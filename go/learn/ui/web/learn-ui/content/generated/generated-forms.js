(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = ContentGenerated.enums;
    var cEnums = ContentCurriculum.enums;

    ContentGenerated.forms = {
        GeneratedLesson: f.form('Generated Lesson', [
            f.section('Lesson Info', [
                ...f.text('title', 'Title'),
                ...f.text('topic', 'Topic'),
                ...f.text('theme', 'Theme'),
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.select('subject', 'Subject', cEnums.SUBJECT),
                ...f.select('difficulty', 'Difficulty', cEnums.DIFFICULTY),
                ...f.select('status', 'Status', enums.LESSON_STATUS),
                ...f.textarea('objective', 'Objective'),
                ...f.number('estimatedMinutes', 'Estimated Minutes'),
                ...f.textarea('parentInstructions', 'Parent Instructions')
            ]),
            f.section('Materials', [
                ...f.text('materialsNeeded', 'Materials Needed')
            ]),
            f.section('Results', [
                ...f.number('questionsCorrect', 'Questions Correct'),
                ...f.number('questionsTotal', 'Questions Total'),
                ...f.number('actualMinutes', 'Actual Minutes'),
                ...f.textarea('aiObservation', 'AI Observation'),
                ...f.text('onStruggleStrategy', 'On Struggle Strategy')
            ])
        ])
    };
})();
