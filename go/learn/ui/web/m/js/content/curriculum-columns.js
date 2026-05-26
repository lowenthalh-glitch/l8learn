(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = MobileContentCurriculum.render;

    MobileContentCurriculum.columns = {
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

    MobileContentCurriculum.columns.Course[1].primary = true;
    MobileContentCurriculum.columns.Course[2].secondary = true;
    MobileContentCurriculum.columns.Unit[1].primary = true;
    MobileContentCurriculum.columns.Unit[4].secondary = true;
    MobileContentCurriculum.columns.Lesson[1].primary = true;
    MobileContentCurriculum.columns.Lesson[5].secondary = true;
    MobileContentCurriculum.columns.Activity[1].primary = true;
    MobileContentCurriculum.columns.Activity[2].secondary = true;
    MobileContentCurriculum.columns.Worksheet[1].primary = true;
    MobileContentCurriculum.columns.Worksheet[2].secondary = true;
})();
