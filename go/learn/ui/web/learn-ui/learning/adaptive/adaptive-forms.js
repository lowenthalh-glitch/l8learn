(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = LearningAdaptive.enums;

    LearningAdaptive.forms = {
        Skill: f.form('Skill', [
            f.section('Skill Info', [
                ...f.text('name', 'Name', true),
                ...f.textarea('description', 'Description'),
                ...f.select('subject', 'Subject', enums.SUBJECT, true),
                ...f.select('gradeLevel', 'Grade Level', enums.MASTERY_LEVEL),
                ...f.text('domain', 'Domain'),
                ...f.text('subdomain', 'Subdomain'),
                ...f.number('typicalMasteryMinutes', 'Typical Mastery Minutes')
            ])
        ]),
        SkillMastery: f.form('Skill Mastery', [
            f.section('Mastery Info', [
                ...f.reference('studentId', 'Student', 'Student', true),
                ...f.reference('skillId', 'Skill', 'Skill', true),
                ...f.select('level', 'Level', enums.MASTERY_LEVEL),
                ...f.number('confidence', 'Confidence'),
                ...f.number('attemptsCount', 'Attempts'),
                ...f.number('correctCount', 'Correct'),
                ...f.number('currentAccuracy', 'Accuracy'),
                ...f.date('firstAttempted', 'First Attempted'),
                ...f.date('lastAttempted', 'Last Attempted'),
                ...f.date('masteredDate', 'Mastered Date'),
                ...f.number('totalTimeSeconds', 'Total Time (s)')
            ])
        ]),
        LearningPath: f.form('Learning Path', [
            f.section('Path Info', [
                ...f.reference('studentId', 'Student', 'Student', true),
                ...f.select('subject', 'Subject', enums.SUBJECT, true),
                ...f.select('status', 'Status', enums.PATH_STATUS),
                ...f.number('activitiesCompleted', 'Activities Completed'),
                ...f.number('skillsMastered', 'Skills Mastered'),
                ...f.number('totalTimeSeconds', 'Total Time (s)'),
                ...f.date('startedDate', 'Started'),
                ...f.date('lastActive', 'Last Active')
            ])
        ]),
        AdaptationRule: f.form('Adaptation Rule', [
            f.section('Rule Info', [
                ...f.text('name', 'Name', true),
                ...f.textarea('description', 'Description'),
                ...f.select('status', 'Status', enums.RULE_STATUS),
                ...f.number('priority', 'Priority'),
                ...f.select('subjectFilter', 'Subject Filter', enums.SUBJECT),
                ...f.select('trigger', 'Trigger', enums.ADAPT_TRIGGER, true),
                ...f.number('triggerThreshold', 'Trigger Threshold'),
                ...f.number('triggerWindow', 'Trigger Window'),
                ...f.select('strategy', 'Strategy', enums.ADAPT_STRATEGY, true),
                ...f.number('maxApplicationsPerSession', 'Max Per Session'),
                ...f.number('cooldownSeconds', 'Cooldown (s)')
            ])
        ])
    };
})();
