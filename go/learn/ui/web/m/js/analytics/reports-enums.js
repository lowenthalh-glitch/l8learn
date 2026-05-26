(function() {
    'use strict';
    window.MobileReportsAnalytics = window.MobileReportsAnalytics || {};

    MobileReportsAnalytics.enums = {
        REPORT_PERIOD: { 0: 'Unknown', 1: 'Daily', 2: 'Weekly', 3: 'Monthly', 4: 'Quarterly', 5: 'Yearly' },
        ENGAGEMENT_LEVEL: { 0: 'Unknown', 1: 'Disengaged', 2: 'Low', 3: 'Moderate', 4: 'High', 5: 'Exceptional' },
        ENGAGEMENT_LEVEL_CLASSES: { 1: 'layer8d-status-terminated', 2: 'layer8d-status-inactive', 3: 'layer8d-status-pending', 4: 'layer8d-status-active', 5: 'layer8d-status-active' },
        SUBJECT: { 0: 'Unknown', 1: 'Math', 2: 'Reading', 3: 'Science', 4: 'Writing', 5: 'Social Studies' }
    };

    var enums = MobileReportsAnalytics.enums;

    MobileReportsAnalytics.render = {
        reportPeriod: function(v) { return enums.REPORT_PERIOD[v] || v; },
        engagementLevel: Layer8MRenderers.createStatusRenderer(enums.ENGAGEMENT_LEVEL, enums.ENGAGEMENT_LEVEL_CLASSES),
        subject: function(v) { return enums.SUBJECT[v] || v; }
    };

    MobileReportsAnalytics.primaryKeys = {
        ProgressReport: 'reportId',
        EngagementMetric: 'metricId'
    };
})();
