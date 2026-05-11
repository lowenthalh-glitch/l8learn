(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Collab',
        modules: {
            'groups': mod('Groups', '\uD83E\uDD1D', [
                svc('groups', 'Groups', '\uD83D\uDC65', '/70/Collab', 'CollabGroup'),
                svc('tutoring', 'Tutoring', '\uD83C\uDF93', '/70/TutorPair', 'TutorMatch'),
                svc('challenges', 'Challenges', '\uD83C\uDFC6', '/70/Challenge', 'Challenge')
            ])
        },
        submodules: ['CollabGroups']
    });
})();
