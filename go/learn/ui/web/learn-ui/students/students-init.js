/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
*/
(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Students',
        defaultModule: 'people',
        defaultService: 'students',
        sectionSelector: 'people',
        initializerName: 'initializeStudents',
        requiredNamespaces: ['StudentsPeople']
    });

    // Auto-create user accounts when Student/Guardian/Teacher is added
    var userModels = { 'Student': 'createStudentUser', 'Guardian': 'createGuardianUser', 'Teacher': 'createTeacherUser' };
    var origInit = window.initializeStudents;
    window.initializeStudents = function() {
        if (origInit) origInit();
        var origOpenAdd = window.Students && window.Students._openAddModal;
        console.log('[students-init] origOpenAdd:', typeof origOpenAdd, 'LearnUserProvisioning:', !!window.LearnUserProvisioning, 'L8BatchProcessDesktop:', !!window.L8BatchProcessDesktop);
        if (typeof origOpenAdd !== 'function') return;

        window.Students._openAddModal = function(service) {
            // EvalImport uses custom multi-file upload modal
            if (service.model === 'EvalImport' && window.L8BatchProcessDesktop) {
                L8BatchProcessDesktop.openMultiUpload();
                return;
            }

            var provisionFn = userModels[service.model];
            if (!provisionFn) { origOpenAdd.call(this, service); return; }

            var formDef = Layer8DServiceRegistry.getFormDef('Students', service.model);
            if (!formDef) { origOpenAdd.call(this, service); return; }
            var svcConfig = {
                endpoint: Layer8DConfig.resolveEndpoint(service.endpoint),
                primaryKey: Layer8DServiceRegistry.getPrimaryKey('Students', service.model),
                modelName: service.model
            };
            Layer8DPopup.show({
                title: 'Add ' + formDef.title,
                content: Layer8DForms.generateFormHtml(formDef, {}),
                size: 'large',
                showFooter: true,
                saveButtonText: 'Save',
                onSave: async function() {
                    var data = Layer8DForms.collectFormData(formDef);
                    var errors = Layer8DForms.validateFormData(formDef, data);
                    if (errors.length > 0) {
                        Layer8DNotification.error('Validation failed', errors.map(function(e) { return e.message; }));
                        return;
                    }
                    try {
                        var result = await Layer8DForms.saveRecord(svcConfig.endpoint, data, false);
                        Layer8DPopup.close();
                        if (window.Students.refreshCurrentTable) window.Students.refreshCurrentTable();
                        var entity = Object.assign({}, data);
                        if (result && result.id) {
                            entity[svcConfig.primaryKey] = result.id;
                        }
                        LearnUserProvisioning[provisionFn](entity);
                    } catch (err) {
                        Layer8DNotification.error('Error saving', [err.message]);
                    }
                },
                onShow: function(body) {
                    Layer8DForms.setFormContext(formDef, svcConfig);
                    setTimeout(function() { Layer8DForms.attachDatePickers(body); }, 50);
                }
            });
        };

        // Override row click for EvalImport and StudentProfile
        if (Students._showDetailsModal) {
            var origShowDetails = Students._showDetailsModal;
            Students._showDetailsModal = function(service, item, itemId) {
                // EvalImport: show custom detail with "Process All" button
                if (service.model === 'EvalImport' && item && item.processingStatus === 5 && window.L8BatchProcessDesktop) {
                    L8BatchProcessDesktop.openEvalDetail(service, item);
                    return;
                }
                // StudentProfile: show standard detail then add schedule buttons
                if (service.model === 'StudentProfile' && item && window.L8ScheduleGenDesktop) {
                    origShowDetails.call(this, service, item, itemId);
                    // After popup renders, inject schedule buttons
                    setTimeout(function() {
                        var body = Layer8DPopup.getBody();
                        if (!body) return;
                        var btnDiv = document.createElement('div');
                        btnDiv.style.cssText = 'padding:12px;border-top:1px solid var(--layer8d-border);display:flex;gap:8px;';
                        var genBtn = document.createElement('button');
                        genBtn.className = 'layer8d-btn layer8d-btn-primary layer8d-btn-small';
                        genBtn.textContent = 'Generate Schedule';
                        genBtn.onclick = function() {
                            Layer8DPopup.close();
                            L8ScheduleGenDesktop.open(item.studentId);
                        };
                        var viewBtn = document.createElement('button');
                        viewBtn.className = 'layer8d-btn layer8d-btn-secondary layer8d-btn-small';
                        viewBtn.textContent = 'View Schedule';
                        viewBtn.onclick = function() {
                            Layer8DPopup.close();
                            L8ScheduleViewDesktop.openByQuery(item.studentId);
                        };
                        btnDiv.appendChild(genBtn);
                        btnDiv.appendChild(viewBtn);
                        body.appendChild(btnDiv);
                    }, 100);
                    return;
                }
                origShowDetails.call(this, service, item, itemId);
            };
        }
    };
})();
