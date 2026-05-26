(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = MobileReportsAnalytics.enums;

    MobileReportsAnalytics.forms = {
        ProgressReport: f.form('Progress Report', [
            f.section('Report Info', [
                ...f.reference('studentId', 'Student', 'Student', true),
                ...f.select('subject', 'Subject', enums.SUBJECT, true),
                ...f.select('period', 'Period', enums.REPORT_PERIOD, true),
                ...f.date('periodStart', 'Period Start'),
                ...f.date('periodEnd', 'Period End'),
                ...f.number('totalTimeMinutes', 'Total Time (min)'),
                ...f.number('sessionsCount', 'Sessions'),
                ...f.number('daysActive', 'Days Active'),
                ...f.number('activitiesCompleted', 'Activities Completed'),
                ...f.number('skillsPracticed', 'Skills Practiced'),
                ...f.number('skillsMastered', 'Skills Mastered'),
                ...f.number('averageScore', 'Average Score'),
                ...f.number('scoreTrend', 'Score Trend'),
                ...f.select('engagement', 'Engagement', enums.ENGAGEMENT_LEVEL),
                ...f.number('streakDays', 'Streak Days'),
                ...f.textarea('aiSummary', 'AI Summary'),
                ...f.textarea('aiRecommendations', 'AI Recommendations')
            ])
        ]),
        EngagementMetric: f.form('Engagement Metric', [
            f.section('Engagement Info', [
                ...f.reference('studentId', 'Student', 'Student', true),
                ...f.select('currentLevel', 'Level', enums.ENGAGEMENT_LEVEL),
                ...f.number('currentStreakDays', 'Current Streak'),
                ...f.number('longestStreakDays', 'Longest Streak'),
                ...f.number('totalPoints', 'Total Points'),
                ...f.number('badgesEarned', 'Badges Earned'),
                ...f.number('avgSessionMinutes', 'Avg Session (min)'),
                ...f.number('avgDailyMinutes', 'Avg Daily (min)'),
                ...f.number('weeklyGoalPercent', 'Weekly Goal %'),
                ...f.date('lastSessionDate', 'Last Session'),
                ...f.number('daysSinceLastSession', 'Days Since Last')
            ])
        ])
    };
})();
