(function() {
    'use strict';
    window.HistoryData = window.HistoryData || {};

    HistoryData.enums = {
        GROWTH_RATING: { 0: 'Unknown', 1: 'Well Below', 2: 'Below', 3: 'Typical', 4: 'Above', 5: 'Well Above' },
        GROWTH_RATING_CLASSES: { 1: 'layer8d-status-terminated', 2: 'layer8d-status-inactive', 3: 'layer8d-status-pending', 4: 'layer8d-status-active', 5: 'layer8d-status-active' },
        RISK_LEVEL: { 0: 'Unknown', 1: 'On Track', 2: 'Watch', 3: 'At Risk', 4: 'Critical' },
        RISK_LEVEL_CLASSES: { 1: 'layer8d-status-active', 2: 'layer8d-status-pending', 3: 'layer8d-status-inactive', 4: 'layer8d-status-terminated' },
        AGGREGATION_LEVEL: { 0: 'Unknown', 1: 'Student', 2: 'Classroom', 3: 'School', 4: 'District' },
        SNAPSHOT_TYPE: { 0: 'Unknown', 1: 'Weekly', 2: 'Monthly', 3: 'Quarterly', 4: 'Semester', 5: 'Year End' },
        SUBJECT: { 0: 'Unknown', 1: 'Math', 2: 'Reading', 3: 'Science', 4: 'Writing', 5: 'Social Studies' },
        MASTERY_LEVEL: { 0: 'Unknown', 1: 'Not Started', 2: 'Emerging', 3: 'Developing', 4: 'Proficient', 5: 'Mastered', 6: 'Exemplary' }
    };

    var enums = HistoryData.enums;

    HistoryData.render = {
        growthRating: Layer8DRenderers.createStatusRenderer(enums.GROWTH_RATING, enums.GROWTH_RATING_CLASSES),
        riskLevel: Layer8DRenderers.createStatusRenderer(enums.RISK_LEVEL, enums.RISK_LEVEL_CLASSES),
        aggregationLevel: function(v) { return enums.AGGREGATION_LEVEL[v] || v; },
        snapshotType: function(v) { return enums.SNAPSHOT_TYPE[v] || v; },
        subject: function(v) { return enums.SUBJECT[v] || v; },
        masteryLevel: function(v) { return enums.MASTERY_LEVEL[v] || v; }
    };

    HistoryData.primaryKeys = {
        GrowthRecord: 'growthId',
        CohortSnapshot: 'snapshotId',
        RiskAssessment: 'assessmentId',
        StandardMastery: 'standardMasteryId',
        ContentEffect: 'effectId'
    };
})();
