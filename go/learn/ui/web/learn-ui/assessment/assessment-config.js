(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Assessment',
        modules: {
            'results': mod('Results', '\uD83D\uDCDD', [
                svc('sessions', 'Sessions', '\uD83D\uDCBB', '/40/LearnSess', 'LearningSession'),
                svc('scores', 'Scores', '\uD83D\uDCCA', '/40/Score', 'Score'),
                svc('benchmarks', 'Benchmarks', '\uD83C\uDFAF', '/40/Benchmark', 'Benchmark'),
                svc('scans', 'Scans', '\uD83D\uDCF7', '/40/WkshtScan', 'WorksheetScan', { readOnly: true })
            ])
        },
        submodules: ['AssessmentResults']
    });
})();
