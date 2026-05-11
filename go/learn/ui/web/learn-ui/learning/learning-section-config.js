(function() {
    'use strict';
    Layer8SectionConfigs.register('learning', {
        title: 'Learning',
        subtitle: 'Adaptive Learning & Mastery',
        icon: '\uD83E\uDDE0',
        initFn: 'initializeLearning',
        modules: [{
            key: 'adaptive',
            label: 'Adaptive',
            icon: '\uD83E\uDDE0',
            isDefault: true,
            services: [
                { key: 'skills', label: 'Skills', icon: '\uD83C\uDFAF', isDefault: true },
                { key: 'mastery', label: 'Mastery', icon: '\uD83C\uDFC6' },
                { key: 'paths', label: 'Paths', icon: '\uD83D\uDDFA\uFE0F' },
                { key: 'rules', label: 'Rules', icon: '\u2699\uFE0F' }
            ]
        }]
    });
})();
