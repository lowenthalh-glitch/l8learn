(function() {
    'use strict';
    window.LearningAdaptive = window.LearningAdaptive || {};

    LearningAdaptive.enums = {
        MASTERY_LEVEL: { 0: 'Unknown', 1: 'Not Started', 2: 'Emerging', 3: 'Developing', 4: 'Proficient', 5: 'Mastered', 6: 'Exemplary' },
        MASTERY_LEVEL_CLASSES: { 1: 'layer8d-status-inactive', 2: 'layer8d-status-pending', 3: 'layer8d-status-pending', 4: 'layer8d-status-active', 5: 'layer8d-status-active', 6: 'layer8d-status-active' },
        PATH_STATUS: { 0: 'Unknown', 1: 'Active', 2: 'Paused', 3: 'Completed' },
        PATH_STATUS_CLASSES: { 1: 'layer8d-status-active', 2: 'layer8d-status-pending', 3: 'layer8d-status-completed' },
        SUBJECT: { 0: 'Unknown', 1: 'Math', 2: 'Reading', 3: 'Science', 4: 'Writing', 5: 'Social Studies' },
        ADAPT_STRATEGY: { 0: 'Unknown', 1: 'Repeat', 2: 'Scaffold', 3: 'Alternate', 4: 'Review', 5: 'Advance', 6: 'Enrich', 7: 'Break' },
        ADAPT_TRIGGER: { 0: 'Unknown', 1: 'Score Below', 2: 'Score Above', 3: 'Streak Correct', 4: 'Streak Incorrect', 5: 'Time Exceeded', 6: 'Time Too Fast', 7: 'Hints Exhausted', 8: 'Engagement Drop', 9: 'Session Duration', 10: 'Mastery Achieved' },
        RULE_STATUS: { 0: 'Unknown', 1: 'Draft', 2: 'Active', 3: 'Disabled' },
        RULE_STATUS_CLASSES: { 1: 'layer8d-status-pending', 2: 'layer8d-status-active', 3: 'layer8d-status-inactive' }
    };

    var enums = LearningAdaptive.enums;

    LearningAdaptive.render = {
        masteryLevel: Layer8DRenderers.createStatusRenderer(enums.MASTERY_LEVEL, enums.MASTERY_LEVEL_CLASSES),
        pathStatus: Layer8DRenderers.createStatusRenderer(enums.PATH_STATUS, enums.PATH_STATUS_CLASSES),
        subject: function(v) { return enums.SUBJECT[v] || v; },
        strategy: function(v) { return enums.ADAPT_STRATEGY[v] || v; },
        trigger: function(v) { return enums.ADAPT_TRIGGER[v] || v; },
        ruleStatus: Layer8DRenderers.createStatusRenderer(enums.RULE_STATUS, enums.RULE_STATUS_CLASSES)
    };

    LearningAdaptive.primaryKeys = {
        Skill: 'skillId',
        SkillMastery: 'masteryId',
        LearningPath: 'pathId',
        AdaptationRule: 'ruleId'
    };
})();
