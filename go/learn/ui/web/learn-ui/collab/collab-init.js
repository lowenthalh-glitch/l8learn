(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Collab',
        defaultModule: 'groups',
        defaultService: 'groups',
        sectionSelector: 'groups',
        initializerName: 'initializeCollab',
        requiredNamespaces: ['CollabGroups']
    });
})();
