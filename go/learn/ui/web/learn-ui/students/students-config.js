/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
*/
(function() {
    'use strict';
    Layer8ModuleConfigFactory.create({
        namespace: 'Students',
        modules: {
            'people': {
                label: 'People', icon: '👨‍🎓',
                services: [
                    { key: 'students', label: 'Students', icon: '👨‍🎓', endpoint: '/20/Student', model: 'Student' },
                    { key: 'profiles', label: 'Profiles', icon: '🧠', endpoint: '/20/Profile', model: 'StudentProfile' },
                    { key: 'guardians', label: 'Guardians', icon: '👨‍👧', endpoint: '/20/Guardian', model: 'Guardian' },
                    { key: 'teachers', label: 'Teachers', icon: '👩‍🏫', endpoint: '/20/Teacher', model: 'Teacher' },
                    { key: 'classrooms', label: 'Classrooms', icon: '🏫', endpoint: '/20/Classroom', model: 'Classroom' },
                    { key: 'schools', label: 'Schools', icon: '🏢', endpoint: '/20/School', model: 'School' },
                    { key: 'districts', label: 'Districts', icon: '🏙️', endpoint: '/20/District', model: 'District' },
                    { key: 'evals', label: 'Evaluations', icon: '📄', endpoint: '/20/EvalImprt', model: 'EvalImport' }
                ]
            }
        },
        submodules: ['StudentsPeople']
    });
})();
