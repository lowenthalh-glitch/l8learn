(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Learning',
        defaultModule: 'adaptive',
        defaultService: 'skills',
        sectionSelector: 'adaptive',
        initializerName: 'initializeLearning',
        requiredNamespaces: ['LearningAdaptive']
    });
})();
