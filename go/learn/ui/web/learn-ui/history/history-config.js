(function() {
    'use strict';
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'History',
        modules: {
            'data': mod('Data', '\uD83D\uDCDA', [
                { key: 'growth', label: 'Growth', icon: '\uD83C\uDF31', endpoint: '/60/Growth', model: 'GrowthRecord', readOnly: true },
                { key: 'cohorts', label: 'Cohorts', icon: '\uD83D\uDC65', endpoint: '/60/Cohort', model: 'CohortSnapshot', readOnly: true },
                { key: 'risk', label: 'Risk', icon: '\u26A0\uFE0F', endpoint: '/60/RiskAssmt', model: 'RiskAssessment', readOnly: true },
                { key: 'standards', label: 'Standards', icon: '\uD83C\uDFAF', endpoint: '/60/StdMastry', model: 'StandardMastery', readOnly: true },
                { key: 'effectiveness', label: 'Effectiveness', icon: '\uD83D\uDCA1', endpoint: '/60/CntEffect', model: 'ContentEffect', readOnly: true }
            ])
        },
        submodules: ['HistoryData']
    });
})();
