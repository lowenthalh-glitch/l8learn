(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = HistoryData.enums;

    HistoryData.forms = {
        GrowthRecord: f.form('Growth Record', [
            f.section('Growth Info', [
                ...f.text('growthId', 'Growth ID', false, { readOnly: true }),
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.select('subject', 'Subject', enums.SUBJECT),
                ...f.text('academicYear', 'Academic Year', false, { readOnly: true }),
                ...f.number('baselineScore', 'Baseline Score', false, { readOnly: true }),
                ...f.number('currentScore', 'Current Score', false, { readOnly: true }),
                ...f.number('absoluteGrowth', 'Absolute Growth', false, { readOnly: true }),
                ...f.number('growthPercentile', 'Growth Percentile', false, { readOnly: true }),
                ...f.select('rating', 'Rating', enums.GROWTH_RATING, false, { readOnly: true }),
                ...f.number('expectedGrowth', 'Expected Growth', false, { readOnly: true }),
                ...f.number('growthVsExpected', 'Growth vs Expected', false, { readOnly: true }),
                ...f.number('totalTimeMinutes', 'Total Time (min)', false, { readOnly: true }),
                ...f.number('totalSessions', 'Total Sessions', false, { readOnly: true })
            ])
        ]),
        CohortSnapshot: f.form('Cohort Snapshot', [
            f.section('Snapshot Info', [
                ...f.text('snapshotId', 'Snapshot ID', false, { readOnly: true }),
                ...f.select('level', 'Level', enums.AGGREGATION_LEVEL, false, { readOnly: true }),
                ...f.text('entityId', 'Entity', false, { readOnly: true }),
                ...f.select('subject', 'Subject', enums.SUBJECT, false, { readOnly: true }),
                ...f.select('type', 'Type', enums.SNAPSHOT_TYPE, false, { readOnly: true }),
                ...f.date('snapshotDate', 'Date'),
                ...f.number('totalStudents', 'Total Students', false, { readOnly: true }),
                ...f.number('activeStudents', 'Active Students', false, { readOnly: true }),
                ...f.number('meanScore', 'Mean Score', false, { readOnly: true }),
                ...f.number('medianScore', 'Median Score', false, { readOnly: true }),
                ...f.number('meanGrowth', 'Mean Growth', false, { readOnly: true }),
                ...f.number('participationRate', 'Participation Rate', false, { readOnly: true })
            ])
        ]),
        RiskAssessment: f.form('Risk Assessment', [
            f.section('Risk Info', [
                ...f.text('assessmentId', 'Assessment ID', false, { readOnly: true }),
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.select('subject', 'Subject', enums.SUBJECT, false, { readOnly: true }),
                ...f.select('riskLevel', 'Risk Level', enums.RISK_LEVEL, false, { readOnly: true }),
                ...f.number('riskScore', 'Risk Score', false, { readOnly: true }),
                ...f.date('assessedDate', 'Assessed Date'),
                ...f.textarea('aiRecommendation', 'AI Recommendation', { readOnly: true }),
                ...f.textarea('teacherNotes', 'Teacher Notes', { readOnly: true }),
                ...f.textarea('interventionPlan', 'Intervention Plan', { readOnly: true })
            ])
        ]),
        StandardMastery: f.form('Standard Mastery', [
            f.section('Standard Info', [
                ...f.text('standardMasteryId', 'ID', false, { readOnly: true }),
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.text('standardId', 'Standard ID', false, { readOnly: true }),
                ...f.text('standardDescription', 'Description', false, { readOnly: true }),
                ...f.select('subject', 'Subject', enums.SUBJECT, false, { readOnly: true }),
                ...f.select('level', 'Level', enums.MASTERY_LEVEL, false, { readOnly: true }),
                ...f.number('score', 'Score', false, { readOnly: true }),
                ...f.number('skillsInStandard', 'Skills in Standard', false, { readOnly: true }),
                ...f.number('skillsMastered', 'Skills Mastered', false, { readOnly: true }),
                ...f.date('lastAssessed', 'Last Assessed')
            ])
        ]),
        ContentEffect: f.form('Content Effectiveness', [
            f.section('Effectiveness Info', [
                ...f.text('effectId', 'Effect ID', false, { readOnly: true }),
                ...f.text('contentName', 'Content Name', false, { readOnly: true }),
                ...f.text('contentType', 'Content Type', false, { readOnly: true }),
                ...f.select('subject', 'Subject', enums.SUBJECT, false, { readOnly: true }),
                ...f.text('academicYear', 'Academic Year', false, { readOnly: true }),
                ...f.number('totalAttempts', 'Total Attempts', false, { readOnly: true }),
                ...f.number('uniqueStudents', 'Unique Students', false, { readOnly: true }),
                ...f.number('meanScore', 'Mean Score', false, { readOnly: true }),
                ...f.number('meanMasteryGain', 'Mastery Gain', false, { readOnly: true }),
                ...f.number('masteryGainPerMinute', 'Gain/Min', false, { readOnly: true }),
                ...f.number('effectivenessPercentile', 'Effectiveness %ile', false, { readOnly: true }),
                ...f.textarea('aiAnalysis', 'AI Analysis', { readOnly: true })
            ])
        ])
    };
})();
