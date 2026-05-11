(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Content',
        defaultModule: 'curriculum',
        defaultService: 'courses',
        sectionSelector: 'curriculum',
        initializerName: 'initializeContent',
        requiredNamespaces: ['ContentCurriculum', 'ContentGenerated']
    });
})();
