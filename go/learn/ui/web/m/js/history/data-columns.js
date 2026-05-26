(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = MobileDataHistory.render;

    MobileDataHistory.columns = {
        GrowthRecord: [
            ...col.id('growthId'),
            ...col.col('studentId', 'Student'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.col('academicYear', 'Year'),
            ...col.number('baselineScore', 'Baseline'),
            ...col.number('currentScore', 'Current'),
            ...col.number('absoluteGrowth', 'Growth'),
            ...col.status('rating', 'Rating', null, render.growthRating)
        ],
        CohortSnapshot: [
            ...col.id('snapshotId'),
            ...col.enum('level', 'Level', null, render.aggregationLevel),
            ...col.col('entityId', 'Entity'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.enum('type', 'Type', null, render.snapshotType),
            ...col.date('snapshotDate', 'Date'),
            ...col.number('totalStudents', 'Students'),
            ...col.number('meanScore', 'Mean Score'),
            ...col.number('meanGrowth', 'Mean Growth')
        ],
        RiskAssessment: [
            ...col.id('assessmentId'),
            ...col.col('studentId', 'Student'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.status('riskLevel', 'Risk', null, render.riskLevel),
            ...col.number('riskScore', 'Score'),
            ...col.date('assessedDate', 'Assessed')
        ],
        StandardMastery: [
            ...col.id('standardMasteryId'),
            ...col.col('studentId', 'Student'),
            ...col.col('standardId', 'Standard'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.enum('level', 'Level', null, render.masteryLevel),
            ...col.number('score', 'Score'),
            ...col.number('skillsMastered', 'Mastered'),
            ...col.date('lastAssessed', 'Last Assessed')
        ],
        ContentEffect: [
            ...col.id('effectId'),
            ...col.col('contentName', 'Content'),
            ...col.col('contentType', 'Type'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.number('totalAttempts', 'Attempts'),
            ...col.number('uniqueStudents', 'Students'),
            ...col.number('meanScore', 'Mean Score'),
            ...col.number('meanMasteryGain', 'Mastery Gain')
        ]
    };

    MobileDataHistory.columns.GrowthRecord[1].primary = true;
    MobileDataHistory.columns.GrowthRecord[7].secondary = true;
    MobileDataHistory.columns.CohortSnapshot[1].primary = true;
    MobileDataHistory.columns.CohortSnapshot[4].secondary = true;
    MobileDataHistory.columns.RiskAssessment[1].primary = true;
    MobileDataHistory.columns.RiskAssessment[3].secondary = true;
    MobileDataHistory.columns.StandardMastery[1].primary = true;
    MobileDataHistory.columns.StandardMastery[4].secondary = true;
    MobileDataHistory.columns.ContentEffect[1].primary = true;
    MobileDataHistory.columns.ContentEffect[3].secondary = true;
})();
