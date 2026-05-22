/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Desktop batch-process wrapper — custom Add modal for EvalImport
and "Process All" button. Uses L8BatchProcess core.
*/
(function() {
    'use strict';

    function getHeaders() {
        return Object.assign({ 'Content-Type': 'application/json' },
            typeof getAuthHeaders === 'function' ? getAuthHeaders() : {});
    }

    function getEndpoint() {
        return Layer8DConfig.resolveEndpoint('/20/EvalImprt');
    }

    // Override Add modal for EvalImport to support multi-file upload
    window.L8BatchProcessDesktop = {

        openMultiUpload: function() {
            var regConf = (typeof Layer8DReferenceRegistry !== 'undefined' && Layer8DReferenceRegistry.Student) || {};
            var refCfg = JSON.stringify({
                modelName: 'Student',
                endpoint: Layer8DConfig.resolveEndpoint('/20/Student'),
                idColumn: regConf.idColumn || 'studentId',
                displayColumn: regConf.displayColumn || 'lastName',
                selectColumns: regConf.selectColumns,
                displayLabel: 'Student'
            });
            var refHtml = '<label>Student</label>' +
                "<input type=\"text\" id=\"l8-batch-student\" name=\"studentId\" " +
                "class=\"reference-input\" " +
                "data-ref-config='" + Layer8DUtils.escapeAttr(refCfg) + "' " +
                "data-ref-id=\"\" " +
                "data-lookup-model=\"Student\" " +
                "data-field-key=\"studentId\" " +
                "readonly placeholder=\"Click to select...\">";
            var content = L8BatchProcess.renderUploadFormHtml(refHtml);

            Layer8DPopup.show({
                title: 'Upload Evaluations',
                content: content,
                size: 'large',
                showFooter: true,
                saveButtonText: 'Upload All',
                onSave: async function() {
                    var body = Layer8DPopup.getBody();
                    var studentInput = body.querySelector('#l8-batch-student');
                    var studentId = studentInput ? studentInput.dataset.refId : '';
                    if (!studentId) {
                        Layer8DNotification.error('Please select a student');
                        return;
                    }
                    var fileInput = body.querySelector('#l8-batch-files');
                    if (!fileInput || !fileInput.files || fileInput.files.length === 0) {
                        Layer8DNotification.error('Please select at least one PDF file');
                        return;
                    }

                    var results = await L8BatchProcess.uploadFiles(
                        fileInput.files, studentId,
                        function(file) { return Layer8FileUpload.upload(file, 'general', '1'); },
                        function(evalData) {
                            return fetch(getEndpoint(), {
                                method: 'POST', headers: getHeaders(),
                                body: JSON.stringify(evalData)
                            }).then(function(r) { if (!r.ok) throw new Error('POST failed'); });
                        },
                        function(cur, total, name) {
                            Layer8DNotification.info('Uploading ' + cur + ' of ' + total + ': ' + name);
                        }
                    );

                    Layer8DPopup.close();
                    if (results.success > 0) {
                        Layer8DNotification.success(results.success + ' file(s) uploaded. Cleaning in progress...');
                    }
                    if (results.failed > 0) {
                        Layer8DNotification.error(results.failed + ' file(s) failed', results.errors);
                    }
                    if (window.Students && Students.refreshCurrentTable) Students.refreshCurrentTable();
                },
                onShow: function(body) {
                    var dropzone = body.querySelector('#l8-batch-dropzone');
                    var fileInput = body.querySelector('#l8-batch-files');
                    var fileList = body.querySelector('#l8-batch-file-list');
                    if (dropzone && fileInput) {
                        dropzone.onclick = function() { fileInput.click(); };
                        dropzone.ondragover = function(e) { e.preventDefault(); dropzone.style.borderColor = 'var(--layer8d-primary)'; };
                        dropzone.ondragleave = function() { dropzone.style.borderColor = 'var(--layer8d-border)'; };
                        dropzone.ondrop = function(e) {
                            e.preventDefault();
                            dropzone.style.borderColor = 'var(--layer8d-border)';
                            fileInput.files = e.dataTransfer.files;
                            fileList.innerHTML = L8BatchProcess.renderFileList(fileInput.files);
                        };
                        fileInput.onchange = function() {
                            fileList.innerHTML = L8BatchProcess.renderFileList(fileInput.files);
                        };
                    }
                    // Attach reference picker for student
                    setTimeout(function() { Layer8DForms.attachReferencePickers(body); }, 50);
                }
            });
        },

        // Trigger batch LLM processing for all CLEANED evals
        processAll: async function(tableData) {
            var result = await L8BatchProcess.triggerProcessAll(
                tableData,
                function(evalRecord) {
                    return fetch(getEndpoint(), {
                        method: 'PUT', headers: getHeaders(),
                        body: JSON.stringify(evalRecord)
                    }).then(function(r) { if (!r.ok) throw new Error('PUT failed'); });
                }
            );
            if (result.triggered) {
                Layer8DNotification.success(result.message);
            } else {
                Layer8DNotification.warning(result.message);
            }
            if (window.Students && Students.refreshCurrentTable) Students.refreshCurrentTable();
        },

        // Open eval detail with "View Document" and "Process All" buttons
        openEvalDetail: function(service, item) {
            var studentId = item.studentId;
            var html = '<div style="margin-bottom:16px;">';
            html += '<p><strong>Status:</strong> ' + (StudentsPeople.enums.EVAL_PROCESSING_STATUS[item.processingStatus] || 'Unknown') + '</p>';
            if (item.filePath) {
                html += '<p><button class="layer8d-btn layer8d-btn-primary layer8d-btn-small" id="l8-view-cleaned-btn">View Cleaned Document</button></p>';
            }
            html += '<hr style="margin:12px 0;">';
            html += '<p style="color:var(--layer8d-text-medium);">When you have reviewed all cleaned documents and confirmed no sensitive data remains, click below to send all evaluations to the AI for analysis.</p>';
            html += '<button class="layer8d-btn layer8d-btn-primary" id="l8-process-all-btn" style="margin-top:8px;">Process All Evaluations for This Student</button>';
            html += '</div>';

            Layer8DPopup.show({
                title: 'Evaluation Review',
                content: html,
                size: 'medium',
                onShow: function(body) {
                    var viewBtn = body.querySelector('#l8-view-cleaned-btn');
                    if (viewBtn) {
                        viewBtn.onclick = function() {
                            LearnFileViewer.view(item.filePath, 'Cleaned Evaluation');
                        };
                    }
                    var processBtn = body.querySelector('#l8-process-all-btn');
                    if (processBtn) {
                        processBtn.onclick = function() {
                            processBtn.disabled = true;
                            processBtn.textContent = 'Processing...';
                            // Fetch all evals for this student and trigger
                            var endpoint = Layer8DConfig.resolveEndpoint('/20/EvalImprt');
                            var query = encodeURIComponent(JSON.stringify({text: 'select * from EvalImport where studentId=' + studentId}));
                            fetch(endpoint + '?body=' + query, { headers: getHeaders() })
                                .then(function(r) { return r.json(); })
                                .then(function(data) {
                                    return L8BatchProcessDesktop.processAll(data.list || []);
                                })
                                .then(function() {
                                    Layer8DPopup.close();
                                });
                        };
                    }
                }
            });
        },

        // Inject "Process All" button into the EvalImport table container
        injectProcessButton: function(containerId, tableData) {
            var container = document.getElementById(containerId);
            if (!container) return;
            // Remove existing button
            var existing = container.querySelector('.l8-process-all-btn');
            if (existing) existing.remove();

            var counts = L8BatchProcess.getStatusCounts(tableData);
            if (counts.cleaned === 0) return;

            var btn = document.createElement('button');
            btn.className = 'l8-process-all-btn layer8d-btn layer8d-btn-primary layer8d-btn-small';
            btn.style.cssText = 'margin:8px 0;';
            btn.textContent = 'Process All Evaluations (' + counts.cleaned + ' ready)';
            btn.onclick = function() {
                L8BatchProcessDesktop.processAll(tableData);
            };
            container.insertBefore(btn, container.firstChild);
        }
    };
})();
