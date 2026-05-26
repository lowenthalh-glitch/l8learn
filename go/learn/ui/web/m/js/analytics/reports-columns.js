(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = MobileReportsAnalytics.render;

    MobileReportsAnalytics.columns = {
        ProgressReport: [
            ...col.id('reportId'),
            ...col.col('studentId', 'Student'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.enum('period', 'Period', null, render.reportPeriod),
            ...col.date('periodStart', 'Start'),
            ...col.date('periodEnd', 'End'),
            ...col.number('totalTimeMinutes', 'Time (min)'),
            ...col.number('skillsMastered', 'Mastered'),
            ...col.number('averageScore', 'Avg Score'),
            ...col.status('engagement', 'Engagement', null, render.engagementLevel)
        ],
        EngagementMetric: [
            ...col.id('metricId'),
            ...col.col('studentId', 'Student'),
            ...col.status('currentLevel', 'Level', null, render.engagementLevel),
            ...col.number('currentStreakDays', 'Streak'),
            ...col.number('totalPoints', 'Points'),
            ...col.number('badgesEarned', 'Badges'),
            ...col.number('avgSessionMinutes', 'Avg Session'),
            ...col.number('weeklyGoalPercent', 'Goal %'),
            ...col.date('lastSessionDate', 'Last Session')
        ]
    };

    MobileReportsAnalytics.columns.ProgressReport[1].primary = true;
    MobileReportsAnalytics.columns.ProgressReport[9].secondary = true;
    MobileReportsAnalytics.columns.EngagementMetric[1].primary = true;
    MobileReportsAnalytics.columns.EngagementMetric[2].secondary = true;
})();
