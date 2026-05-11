(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = LearningAdaptive.render;

    LearningAdaptive.columns = {
        Skill: [
            ...col.id('skillId'),
            ...col.col('name', 'Name'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.col('gradeLevel', 'Grade'),
            ...col.col('domain', 'Domain')
        ],
        SkillMastery: [
            ...col.id('masteryId'),
            ...col.col('studentId', 'Student'),
            ...col.col('skillId', 'Skill'),
            ...col.status('level', 'Level', null, render.masteryLevel),
            ...col.number('currentAccuracy', 'Accuracy'),
            ...col.number('attemptsCount', 'Attempts')
        ],
        LearningPath: [
            ...col.id('pathId'),
            ...col.col('studentId', 'Student'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.status('status', 'Status', null, render.pathStatus),
            ...col.number('activitiesCompleted', 'Activities'),
            ...col.number('skillsMastered', 'Mastered')
        ],
        AdaptationRule: [
            ...col.id('ruleId'),
            ...col.col('name', 'Name'),
            ...col.status('status', 'Status', null, render.ruleStatus),
            ...col.enum('trigger', 'Trigger', null, render.trigger),
            ...col.enum('strategy', 'Strategy', null, render.strategy),
            ...col.number('priority', 'Priority')
        ]
    };
})();
