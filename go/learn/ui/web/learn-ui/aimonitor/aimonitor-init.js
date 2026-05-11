(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'AIMonitor',
        defaultModule: 'data',
        defaultService: 'promptlogs',
        sectionSelector: 'data',
        initializerName: 'initializeAIMonitor',
        requiredNamespaces: ['AIMonitorData']
    });
})();
