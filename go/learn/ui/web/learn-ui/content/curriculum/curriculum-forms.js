(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = ContentCurriculum.enums;

    ContentCurriculum.forms = {
        Course: f.form('Course', [
            f.section('Course Info', [
                ...f.text('name', 'Name', true),
                ...f.select('subject', 'Subject', enums.SUBJECT, true),
                ...f.select('minGrade', 'Min Grade', enums.GRADE),
                ...f.select('maxGrade', 'Max Grade', enums.GRADE),
                ...f.select('status', 'Status', enums.STATUS),
                ...f.number('estimatedHours', 'Estimated Hours'),
                ...f.textarea('description', 'Description')
            ])
        ]),
        Unit: f.form('Unit', [
            f.section('Unit Info', [
                ...f.text('name', 'Name', true),
                ...f.reference('courseId', 'Course', 'Course', true),
                ...f.number('sequenceOrder', 'Sequence Order'),
                ...f.select('status', 'Status', enums.STATUS),
                ...f.number('estimatedMinutes', 'Estimated Minutes')
            ])
        ]),
        Lesson: f.form('Lesson', [
            f.section('Lesson Info', [
                ...f.text('name', 'Name', true),
                ...f.reference('unitId', 'Unit', 'Unit', true),
                ...f.number('sequenceOrder', 'Sequence Order'),
                ...f.select('difficulty', 'Difficulty', enums.DIFFICULTY),
                ...f.select('status', 'Status', enums.STATUS),
                ...f.number('estimatedMinutes', 'Estimated Minutes'),
                ...f.textarea('instructionText', 'Instructions')
            ])
        ]),
        Activity: f.form('Activity', [
            f.section('Activity Info', [
                ...f.text('name', 'Name', true),
                ...f.reference('lessonId', 'Lesson', 'Lesson', true),
                ...f.select('activityType', 'Type', enums.ACTIVITY_TYPE),
                ...f.select('difficulty', 'Difficulty', enums.DIFFICULTY),
                ...f.select('status', 'Status', enums.STATUS),
                ...f.number('pointsPossible', 'Points Possible'),
                ...f.number('estimatedSeconds', 'Est. Seconds'),
                ...f.textarea('instructions', 'Instructions')
            ])
        ]),
        Worksheet: f.form('Worksheet', [
            f.section('Worksheet Info', [
                ...f.text('name', 'Name', true),
                ...f.select('subject', 'Subject', enums.SUBJECT),
                ...f.number('questionCount', 'Question Count'),
                ...f.textarea('instructionsText', 'Instructions')
            ])
        ])
    };
})();
