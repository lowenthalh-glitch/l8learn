/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Schedule view core — renders weekly schedule as a visual table.
*/
(function() {
    'use strict';

    var ACTIVITY_COLORS = {
        'movement_warmup': '#22c55e',
        'academic':        '#3b82f6',
        'therapy':         '#8b5cf6',
        'creative':        '#f59e0b',
        'break':           '#94a3b8',
        'cleanup':         '#64748b'
    };

    var ACTIVITY_ICONS = {
        'movement_warmup': '🏃',
        'academic':        '📚',
        'therapy':         '🗣️',
        'creative':        '🎨',
        'break':           '☕',
        'cleanup':         '🧹'
    };

    function minutesToTime(m) {
        var h = Math.floor(m / 60);
        var min = m % 60;
        var ampm = h >= 12 ? 'PM' : 'AM';
        if (h > 12) h -= 12;
        if (h === 0) h = 12;
        return h + ':' + (min < 10 ? '0' : '') + min + ' ' + ampm;
    }

    window.L8ScheduleView = {

        // Group blocks by day (from blockId prefix: mon-, tue-, wed-, etc.)
        groupByDay: function(blocks) {
            var days = { 'mon': [], 'tue': [], 'wed': [], 'thu': [], 'fri': [] };
            var dayNames = { 'mon': 'Monday', 'tue': 'Tuesday', 'wed': 'Wednesday', 'thu': 'Thursday', 'fri': 'Friday' };
            for (var i = 0; i < (blocks || []).length; i++) {
                var b = blocks[i];
                var prefix = (b.blockId || '').substring(0, 3);
                if (days[prefix]) days[prefix].push(b);
            }
            var result = [];
            var keys = ['mon', 'tue', 'wed', 'thu', 'fri'];
            for (var k = 0; k < keys.length; k++) {
                result.push({ key: keys[k], name: dayNames[keys[k]], blocks: days[keys[k]] });
            }
            return result;
        },

        // Render the full weekly schedule HTML
        renderScheduleHtml: function(schedule, lessons, progressText) {
            var blocks = schedule.blocks || [];
            var days = this.groupByDay(blocks);
            var lessonMap = {};
            for (var i = 0; i < (lessons || []).length; i++) {
                var l = lessons[i];
                if (l.blockId) lessonMap[l.blockId] = l;
            }

            var html = '<div class="l8-schedule-view">';

            // Progress bar
            if (progressText) {
                html += '<div style="padding:8px 12px;margin-bottom:12px;background:var(--layer8d-bg-light);border-radius:6px;font-size:13px;">';
                html += '<strong>' + Layer8DUtils.escapeHtml(progressText) + '</strong>';
                if (schedule.lessonsTotal > 0) {
                    var pct = Math.round((schedule.lessonsGenerated / schedule.lessonsTotal) * 100);
                    html += '<div style="background:var(--layer8d-border);border-radius:4px;height:6px;margin-top:4px;">';
                    html += '<div style="background:var(--layer8d-primary);height:6px;border-radius:4px;width:' + pct + '%;"></div>';
                    html += '</div>';
                }
                html += '</div>';
            }

            // Day columns
            html += '<div style="display:flex;gap:8px;overflow-x:auto;">';
            for (var d = 0; d < days.length; d++) {
                var day = days[d];
                html += '<div style="flex:1;min-width:180px;">';
                html += '<div style="font-weight:bold;padding:8px;background:var(--layer8d-primary);color:#fff;border-radius:6px 6px 0 0;text-align:center;">' + day.name + '</div>';
                html += '<div style="border:1px solid var(--layer8d-border);border-top:none;border-radius:0 0 6px 6px;padding:4px;">';
                for (var b = 0; b < day.blocks.length; b++) {
                    html += this.renderBlockHtml(day.blocks[b], lessonMap[day.blocks[b].blockId]);
                }
                html += '</div></div>';
            }
            html += '</div></div>';
            return html;
        },

        // Render a single block
        renderBlockHtml: function(block, lesson) {
            var color = ACTIVITY_COLORS[block.activityType] || '#94a3b8';
            var icon = ACTIVITY_ICONS[block.activityType] || '📋';
            var time = minutesToTime(block.startMinute);
            var hasLesson = !!lesson;
            var clickAttr = hasLesson ? 'data-lesson-id="' + Layer8DUtils.escapeAttr(lesson.generatedLessonId) + '" style="cursor:pointer;"' : '';

            var html = '<div class="l8-schedule-block" ' + clickAttr +
                ' style="margin:3px 0;padding:6px 8px;border-left:4px solid ' + color + ';background:var(--layer8d-bg-white);border-radius:4px;font-size:12px;">';
            html += '<div style="display:flex;justify-content:space-between;align-items:center;">';
            html += '<span style="font-weight:600;">' + icon + ' ' + time + '</span>';
            html += '<span style="color:var(--layer8d-text-muted);">' + block.durationMinutes + 'min</span>';
            html += '</div>';
            html += '<div style="margin-top:2px;color:var(--layer8d-text-medium);line-height:1.3;">' +
                Layer8DUtils.escapeHtml(block.description || '').substring(0, 80) + '</div>';
            if (hasLesson) {
                html += '<div style="margin-top:3px;font-size:11px;color:var(--layer8d-primary);">📖 ' +
                    Layer8DUtils.escapeHtml(lesson.title) + '</div>';
            }
            html += '</div>';
            return html;
        }
    };
})();
