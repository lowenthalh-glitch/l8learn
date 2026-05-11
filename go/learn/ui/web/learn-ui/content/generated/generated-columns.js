(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = ContentGenerated.render;
    var cRender = ContentCurriculum.render;

    ContentGenerated.columns = {
        GeneratedLesson: [
            ...col.id('generatedLessonId'),
            ...col.col('title', 'Title'),
            ...col.col('topic', 'Topic'),
            ...col.col('theme', 'Theme'),
            ...col.col('studentId', 'Student'),
            ...col.enum('subject', 'Subject', null, cRender.subject),
            ...col.enum('difficulty', 'Difficulty', null, cRender.difficulty),
            ...col.status('status', 'Status', null, render.lessonStatus),
            ...col.number('estimatedMinutes', 'Est. Min'),
            ...col.number('questionsCorrect', 'Correct'),
            ...col.number('questionsTotal', 'Total'),
            ...col.date('generatedAt', 'Generated')
        ]
    };
})();
