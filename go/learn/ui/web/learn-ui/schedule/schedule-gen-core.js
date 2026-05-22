/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Schedule generation core — platform-independent logic.
*/
(function() {
    'use strict';

    window.L8ScheduleGen = {

        // Build the schedule generation form HTML
        renderFormHtml: function(studentRefHtml) {
            return '<div class="l8-schedule-gen-form">' +
                '<div class="form-group">' + studentRefHtml + '</div>' +
                '<div class="form-group">' +
                '<label>Available Hours</label>' +
                '<input type="number" id="l8-sched-hours" value="4" min="1" max="8" style="width:80px;">' +
                '</div>' +
                '<div class="form-group">' +
                '<label>Parent Energy</label>' +
                '<select id="l8-sched-energy">' +
                '<option value="low">Low</option>' +
                '<option value="medium" selected>Medium</option>' +
                '<option value="high">High</option>' +
                '</select>' +
                '</div>' +
                '<div class="form-group">' +
                '<label>Weather</label>' +
                '<select id="l8-sched-weather">' +
                '<option value="sunny" selected>Sunny</option>' +
                '<option value="rainy">Rainy</option>' +
                '<option value="cold">Cold</option>' +
                '</select>' +
                '</div>' +
                '<div class="form-group">' +
                '<label>Appointments (optional)</label>' +
                '<textarea id="l8-sched-appointments" rows="2" placeholder="e.g. dentist at 2pm"></textarea>' +
                '</div>' +
                '</div>';
        },

        // Collect form data from the DOM
        collectFormData: function(body, studentId) {
            return {
                familyId: 'FAM-' + studentId,
                scheduleDate: Math.floor(Date.now() / 1000),
                availableHours: parseInt(body.querySelector('#l8-sched-hours').value) || 4,
                parentEnergy: body.querySelector('#l8-sched-energy').value,
                weather: body.querySelector('#l8-sched-weather').value,
                appointments: (body.querySelector('#l8-sched-appointments').value || '').split('\n').filter(function(a) { return a.trim(); }),
                customFields: { studentId: studentId }
            };
        },

        // Check if schedule is still generating (for polling)
        isGenerating: function(schedule) {
            return schedule && schedule.lessonsTotal > 0 && schedule.lessonsGenerated < schedule.lessonsTotal;
        },

        // Get progress text
        progressText: function(schedule) {
            if (!schedule) return '';
            if (schedule.lessonsTotal === 0) return 'Generating schedule blocks...';
            if (schedule.lessonsGenerated >= schedule.lessonsTotal) {
                return 'Complete — ' + schedule.lessonsTotal + ' lessons ready';
            }
            return 'Generating lesson ' + schedule.lessonsGenerated + ' of ' + schedule.lessonsTotal + '...';
        }
    };
})();
