(function() {
    'use strict';
    Layer8SectionConfigs.register('history', {
        title: 'History',
        subtitle: 'Growth, Cohorts & Risk',
        icon: '\uD83D\uDCDA',
        initFn: 'initializeHistory',
        modules: [{
            key: 'data',
            label: 'Data',
            icon: '\uD83D\uDCDA',
            isDefault: true,
            services: [
                { key: 'growth', label: 'Growth', icon: '\uD83C\uDF31', isDefault: true },
                { key: 'cohorts', label: 'Cohorts', icon: '\uD83D\uDC65' },
                { key: 'risk', label: 'Risk', icon: '\u26A0\uFE0F' },
                { key: 'standards', label: 'Standards', icon: '\uD83C\uDFAF' },
                { key: 'effectiveness', label: 'Effectiveness', icon: '\uD83D\uDCA1' }
            ]
        }]
    });
})();
