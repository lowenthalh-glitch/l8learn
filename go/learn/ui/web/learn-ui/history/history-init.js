(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'History',
        defaultModule: 'data',
        defaultService: 'growth',
        sectionSelector: 'data',
        initializerName: 'initializeHistory',
        requiredNamespaces: ['HistoryData']
    });
})();
