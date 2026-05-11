(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'History',
        modules: {
            'data': mod('Data', '\uD83D\uDCDA', [
                svc('growth', 'Growth', '\uD83C\uDF31', '/60/Growth', 'GrowthRecord', { readOnly: true }),
                svc('cohorts', 'Cohorts', '\uD83D\uDC65', '/60/Cohort', 'CohortSnapshot', { readOnly: true }),
                svc('risk', 'Risk', '\u26A0\uFE0F', '/60/RiskAssmt', 'RiskAssessment', { readOnly: true }),
                svc('standards', 'Standards', '\uD83C\uDFAF', '/60/StdMastry', 'StandardMastery', { readOnly: true }),
                svc('effectiveness', 'Effectiveness', '\uD83D\uDCA1', '/60/CntEffect', 'ContentEffect', { readOnly: true })
            ])
        },
        submodules: ['HistoryData']
    });
})();
