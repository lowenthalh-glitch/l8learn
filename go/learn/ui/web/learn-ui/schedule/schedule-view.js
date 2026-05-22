/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Desktop schedule view wrapper — shows weekly schedule in a popup.
*/
(function() {
    'use strict';

    function getHeaders() {
        return Object.assign({ 'Content-Type': 'application/json' },
            typeof getAuthHeaders === 'function' ? getAuthHeaders() : {});
    }

    var pollTimer = null;

    window.L8ScheduleViewDesktop = {

        // Open schedule view by fetching the latest schedule for a student
        openByQuery: function(studentId) {
            var endpoint = Layer8DConfig.resolveEndpoint('/30/Schedule');
            var query = encodeURIComponent(JSON.stringify({text: 'select * from DailySchedule'}));
            fetch(endpoint + '?body=' + query, { headers: getHeaders() })
                .then(function(r) { return r.json(); })
                .then(function(data) {
                    var schedules = (data.list || []).filter(function(s) {
                        return s.customFields && s.customFields.studentId === studentId;
                    });
                    if (schedules.length === 0) {
                        Layer8DNotification.warning('No schedule found for this student');
                        return;
                    }
                    // Use the most recent schedule
                    var schedule = schedules[schedules.length - 1];
                    L8ScheduleViewDesktop.open(schedule);
                })
                .catch(function(err) {
                    Layer8DNotification.error('Failed to load schedule: ' + err.message);
                });
        },

        // Open the schedule view popup
        open: function(schedule) {
            var self = this;

            // Fetch lessons for this schedule
            var lessonEndpoint = Layer8DConfig.resolveEndpoint('/10/GenLesson');
            var lQuery = encodeURIComponent(JSON.stringify({text: 'select * from GeneratedLesson where scheduleId=' + schedule.scheduleId}));
            fetch(lessonEndpoint + '?body=' + lQuery, { headers: getHeaders() })
                .then(function(r) { return r.json(); })
                .then(function(lessonData) {
                    var lessons = lessonData.list || [];
                    var progress = L8ScheduleGen.progressText(schedule);
                    var html = L8ScheduleView.renderScheduleHtml(schedule, lessons, progress);

                    Layer8DPopup.show({
                        title: 'Weekly Schedule',
                        content: html,
                        size: 'xlarge',
                        onShow: function(body) {
                            // Click handler for lesson blocks
                            body.addEventListener('click', function(e) {
                                var block = e.target.closest('[data-lesson-id]');
                                if (!block) return;
                                var lessonId = block.dataset.lessonId;
                                var lesson = lessons.find(function(l) { return l.generatedLessonId === lessonId; });
                                if (lesson) L8LessonPlayerDesktop.open(lesson);
                            });

                            // Poll for progress if still generating
                            if (L8ScheduleGen.isGenerating(schedule)) {
                                self.startPolling(schedule, lessons);
                            }
                        }
                    });
                });
        },

        startPolling: function(schedule, existingLessons) {
            if (pollTimer) clearInterval(pollTimer);
            pollTimer = setInterval(function() {
                var endpoint = Layer8DConfig.resolveEndpoint('/30/Schedule');
                var query = encodeURIComponent(JSON.stringify({text: 'select * from DailySchedule where scheduleId=' + schedule.scheduleId}));
                fetch(endpoint + '?body=' + query, { headers: getHeaders() })
                    .then(function(r) { return r.json(); })
                    .then(function(data) {
                        var updated = (data.list || [])[0];
                        if (!updated) return;
                        if (!L8ScheduleGen.isGenerating(updated)) {
                            clearInterval(pollTimer);
                            pollTimer = null;
                            Layer8DNotification.success('Schedule generation complete!');
                            Layer8DPopup.close();
                            L8ScheduleViewDesktop.open(updated);
                        }
                    });
            }, 5000);
        }
    };
})();
