(function() {
    'use strict';
    Layer8SectionConfigs.register('assessment', {
        title: 'Assessment',
        subtitle: 'Sessions, Scores & Benchmarks',
        icon: '\uD83D\uDCDD',
        initFn: 'initializeAssessment',
        modules: [{
            key: 'results',
            label: 'Results',
            icon: '\uD83D\uDCDD',
            isDefault: true,
            services: [
                { key: 'sessions', label: 'Sessions', icon: '\uD83D\uDCBB', isDefault: true },
                { key: 'scores', label: 'Scores', icon: '\uD83D\uDCCA' },
                { key: 'benchmarks', label: 'Benchmarks', icon: '\uD83C\uDFAF' },
                { key: 'scans', label: 'Scans', icon: '\uD83D\uDCF7' }
            ]
        }]
    });
})();
