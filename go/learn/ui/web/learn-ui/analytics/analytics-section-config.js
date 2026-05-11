(function() {
    'use strict';
    Layer8SectionConfigs.register('analytics', {
        title: 'Analytics',
        subtitle: 'Progress & Engagement Reports',
        icon: '\uD83D\uDCC8',
        initFn: 'initializeAnalytics',
        modules: [{
            key: 'reports',
            label: 'Reports',
            icon: '\uD83D\uDCC8',
            isDefault: true,
            services: [
                { key: 'progress', label: 'Progress', icon: '\uD83D\uDCCA', isDefault: true },
                { key: 'engagement', label: 'Engagement', icon: '\uD83D\uDD25' }
            ]
        }]
    });
})();
