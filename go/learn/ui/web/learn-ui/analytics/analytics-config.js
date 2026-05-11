(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Analytics',
        modules: {
            'reports': mod('Reports', '\uD83D\uDCC8', [
                svc('progress', 'Progress', '\uD83D\uDCCA', '/50/Progress', 'ProgressReport'),
                svc('engagement', 'Engagement', '\uD83D\uDD25', '/50/Engage', 'EngagementMetric')
            ])
        },
        submodules: ['AnalyticsReports']
    });
})();
