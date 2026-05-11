(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Learning',
        modules: {
            'adaptive': mod('Adaptive', '\uD83E\uDDE0', [
                svc('skills', 'Skills', '\uD83C\uDFAF', '/30/Skill', 'Skill'),
                svc('mastery', 'Mastery', '\uD83C\uDFC6', '/30/Mastery', 'SkillMastery'),
                svc('paths', 'Paths', '\uD83D\uDDFA\uFE0F', '/30/LearnPath', 'LearningPath'),
                svc('rules', 'Rules', '\u2699\uFE0F', '/30/AdaptRule', 'AdaptationRule')
            ])
        },
        submodules: ['LearningAdaptive']
    });
})();
