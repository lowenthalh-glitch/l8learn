(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = ContentCurriculum.render;
    var enums = ContentCurriculum.enums;

    ContentCurriculum.columns = {
        Course: [
            ...col.id('courseId'),
            ...col.col('name', 'Name'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.enum('minGrade', 'Min Grade', null, render.grade),
            ...col.enum('maxGrade', 'Max Grade', null, render.grade),
            ...col.status('status', 'Status', null, render.status),
            ...col.number('estimatedHours', 'Est. Hours')
        ],
        Unit: [
            ...col.id('unitId'),
            ...col.col('name', 'Name'),
            ...col.col('courseId', 'Course'),
            ...col.number('sequenceOrder', 'Order'),
            ...col.status('status', 'Status', null, render.status),
            ...col.number('estimatedMinutes', 'Est. Min')
        ],
        Lesson: [
            ...col.id('lessonId'),
            ...col.col('name', 'Name'),
            ...col.col('unitId', 'Unit'),
            ...col.number('sequenceOrder', 'Order'),
            ...col.enum('difficulty', 'Difficulty', null, render.difficulty),
            ...col.status('status', 'Status', null, render.status)
        ],
        Activity: [
            ...col.id('activityId'),
            ...col.col('name', 'Name'),
            ...col.enum('activityType', 'Type', null, render.activityType),
            ...col.enum('difficulty', 'Difficulty', null, render.difficulty),
            ...col.status('status', 'Status', null, render.status),
            ...col.number('pointsPossible', 'Points')
        ],
        Worksheet: [
            ...col.id('worksheetId'),
            ...col.col('name', 'Name'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.number('questionCount', 'Questions')
        ]
    };
})();
