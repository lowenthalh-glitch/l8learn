/* © 2025 Sharon Aicler (saichler@gmail.com) Layer 8 Ecosystem - Apache 2.0 */
(function() {
    'use strict';
    Layer8SectionConfigs.register('students', {
        title: 'Students',
        subtitle: 'People & Organizations',
        icon: '👨‍🎓',
        initFn: 'initializeStudents',
        modules: [{
            key: 'people', label: 'People', icon: '👨‍🎓', isDefault: true,
            services: [
                { key: 'students', label: 'Students', icon: '👨‍🎓', isDefault: true },
                { key: 'guardians', label: 'Guardians', icon: '👨‍👧' },
                { key: 'teachers', label: 'Teachers', icon: '👩‍🏫' },
                { key: 'classrooms', label: 'Classrooms', icon: '🏫' },
                { key: 'schools', label: 'Schools', icon: '🏢' },
                { key: 'districts', label: 'Districts', icon: '🏙️' }
            ]
        }]
    });
})();
