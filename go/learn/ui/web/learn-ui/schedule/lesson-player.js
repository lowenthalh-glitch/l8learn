/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Desktop lesson player wrapper — shows a generated lesson in a popup.
*/
(function() {
    'use strict';

    window.L8LessonPlayerDesktop = {

        open: function(lesson) {
            var html = L8LessonPlayer.renderLessonHtml(lesson);

            Layer8DPopup.show({
                title: lesson.title || 'Lesson',
                content: html,
                size: 'xlarge'
            });
        }
    };
})();
