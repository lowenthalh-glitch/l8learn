(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Assessment',
        defaultModule: 'results',
        defaultService: 'sessions',
        sectionSelector: 'results',
        initializerName: 'initializeAssessment',
        requiredNamespaces: ['AssessmentResults']
    });
})();
