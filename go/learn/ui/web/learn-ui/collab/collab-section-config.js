(function() {
    'use strict';
    Layer8SectionConfigs.register('collab', {
        title: 'Collaboration',
        subtitle: 'Groups, Tutoring & Challenges',
        icon: '\uD83E\uDD1D',
        initFn: 'initializeCollab',
        modules: [{
            key: 'groups',
            label: 'Groups',
            icon: '\uD83E\uDD1D',
            isDefault: true,
            services: [
                { key: 'groups', label: 'Groups', icon: '\uD83D\uDC65', isDefault: true },
                { key: 'tutoring', label: 'Tutoring', icon: '\uD83C\uDF93' },
                { key: 'challenges', label: 'Challenges', icon: '\uD83C\uDFC6' }
            ]
        }]
    });
})();
