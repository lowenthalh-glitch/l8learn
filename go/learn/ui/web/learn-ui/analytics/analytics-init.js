(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Analytics',
        defaultModule: 'reports',
        defaultService: 'progress',
        sectionSelector: 'reports',
        initializerName: 'initializeAnalytics',
        requiredNamespaces: ['AnalyticsReports']
    });
})();
