/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Lesson player core — renders a generated lesson with steps, questions, and materials.
*/
(function() {
    'use strict';

    var STEP_ICONS = {
        'physical':   '🏃',
        'screen':     '💻',
        'discussion': '💬',
        'worksheet':  '📝',
        'break':      '☕'
    };

    window.L8LessonPlayer = {

        // Render full lesson detail HTML
        renderLessonHtml: function(lesson) {
            if (!lesson) return '<p>Lesson not found</p>';
            var esc = Layer8DUtils.escapeHtml;

            var html = '<div class="l8-lesson-player">';

            // Header
            html += '<div style="margin-bottom:16px;">';
            html += '<h3 style="margin:0 0 4px;">' + esc(lesson.title) + '</h3>';
            html += '<p style="margin:0;color:var(--layer8d-text-medium);">' + esc(lesson.objective) + '</p>';
            if (lesson.theme) {
                html += '<p style="margin:4px 0 0;font-size:13px;color:var(--layer8d-text-muted);">Theme: ' + esc(lesson.theme) + '</p>';
            }
            html += '<div style="margin-top:8px;display:flex;gap:12px;font-size:13px;">';
            html += '<span>⏱️ ' + (lesson.estimatedMinutes || '?') + ' min</span>';
            html += '<span>✅ Pass: ' + (lesson.minCorrectToPass || 0) + '</span>';
            html += '<span>⬆️ Advance: ' + (lesson.minCorrectToAdvance || 0) + '</span>';
            html += '</div>';
            html += '</div>';

            // Materials
            if (lesson.materialsNeeded && lesson.materialsNeeded.length > 0) {
                html += '<div style="margin-bottom:16px;padding:10px;background:var(--layer8d-bg-light);border-radius:6px;">';
                html += '<strong>📦 Materials Needed:</strong><ul style="margin:4px 0 0;padding-left:20px;">';
                for (var m = 0; m < lesson.materialsNeeded.length; m++) {
                    html += '<li>' + esc(lesson.materialsNeeded[m]) + '</li>';
                }
                html += '</ul></div>';
            }

            // Parent instructions
            if (lesson.parentInstructions) {
                html += '<div style="margin-bottom:16px;padding:10px;background:#fef3c7;border-radius:6px;border-left:4px solid #f59e0b;">';
                html += '<strong>👨‍👧 Parent Instructions:</strong><br>' + esc(lesson.parentInstructions);
                html += '</div>';
            }

            // Steps
            var steps = lesson.steps || [];
            html += '<div style="margin-bottom:16px;">';
            html += '<h4 style="margin:0 0 8px;">Lesson Steps (' + steps.length + ')</h4>';
            for (var s = 0; s < steps.length; s++) {
                html += this.renderStepHtml(steps[s]);
            }
            html += '</div>';

            // Struggle strategy
            if (lesson.onStruggleStrategy) {
                html += '<div style="padding:8px;background:var(--layer8d-bg-light);border-radius:6px;font-size:13px;">';
                html += '<strong>💡 If struggling:</strong> ' + esc(lesson.onStruggleStrategy);
                html += '</div>';
            }

            html += '</div>';
            return html;
        },

        // Render a single step
        renderStepHtml: function(step) {
            var esc = Layer8DUtils.escapeHtml;
            var icon = STEP_ICONS[step.stepType] || '📋';

            var html = '<div style="margin:8px 0;padding:10px;border:1px solid var(--layer8d-border);border-radius:6px;">';
            html += '<div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:6px;">';
            html += '<span style="font-weight:600;">' + icon + ' Step ' + step.stepNumber + ': ' + esc(step.title) + '</span>';
            html += '<span style="font-size:12px;color:var(--layer8d-text-muted);">' + step.durationMinutes + ' min | ' + esc(step.parentRole || '') + '</span>';
            html += '</div>';
            html += '<p style="margin:0 0 6px;font-size:13px;line-height:1.5;">' + esc(step.instructions) + '</p>';

            if (step.materialsInstructions) {
                html += '<p style="margin:0 0 6px;font-size:12px;color:var(--layer8d-text-muted);">Materials: ' + esc(step.materialsInstructions) + '</p>';
            }

            // Questions
            var questions = step.questions || [];
            if (questions.length > 0) {
                html += '<div style="margin-top:8px;padding-top:6px;border-top:1px solid var(--layer8d-border);">';
                html += '<strong style="font-size:12px;">Questions (' + questions.length + '):</strong>';
                for (var q = 0; q < questions.length; q++) {
                    html += this.renderQuestionHtml(questions[q], q + 1);
                }
                html += '</div>';
            }

            html += '</div>';
            return html;
        },

        // Render a single question
        renderQuestionHtml: function(question, num) {
            var esc = Layer8DUtils.escapeHtml;
            var html = '<div style="margin:6px 0;padding:6px;background:var(--layer8d-bg-light);border-radius:4px;font-size:13px;">';
            html += '<div><strong>Q' + num + ':</strong> ' + esc(question.prompt) + '</div>';

            // Options
            var options = question.options || [];
            if (options.length > 0) {
                html += '<div style="margin:4px 0 0 16px;">';
                for (var o = 0; o < options.length; o++) {
                    var opt = options[o];
                    var marker = opt.isCorrect ? '✅' : '⬜';
                    html += '<div>' + marker + ' ' + esc(opt.text);
                    if (opt.feedback) html += ' <span style="color:var(--layer8d-text-muted);">(' + esc(opt.feedback) + ')</span>';
                    html += '</div>';
                }
                html += '</div>';
            }

            // Hints
            if (question.hints && question.hints.length > 0) {
                html += '<div style="margin-top:4px;font-size:12px;color:var(--layer8d-text-muted);">Hints: ' +
                    question.hints.map(function(h) { return esc(h); }).join(' → ') + '</div>';
            }

            html += '</div>';
            return html;
        }
    };
})();
