/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Shared finding review logic — platform-independent.
Used by desktop (people-eval-review.js) and mobile (people-eval-review-m.js).
*/
(function() {
    'use strict';

    window.L8EvalReview = {

        renderFindingsHtml: function(evalImport) {
            var findings = evalImport.findings || [];
            var contradictions = evalImport.contradictions || [];
            var html = '';

            html += '<div class="l8-eval-review">';
            html += '<div class="l8-eval-header">';
            html += '<strong>Document Type:</strong> ' + (StudentsPeople.render.evalDocType(evalImport.documentType) || 'Unknown');
            html += ' &nbsp;|&nbsp; <strong>Professional:</strong> ' + Layer8DUtils.escapeHtml(evalImport.professionalName || '-');
            html += '</div>';

            if (evalImport.processingStatus === 4) {
                html += '<div class="l8-eval-error" style="color:var(--layer8d-error);padding:8px;margin:8px 0;border:1px solid var(--layer8d-error);border-radius:4px;">';
                html += '<strong>Processing Failed:</strong> ' + Layer8DUtils.escapeHtml(evalImport.errorMessage || 'Unknown error');
                html += '</div>';
            }

            if (evalImport.processingStatus === 1 || evalImport.processingStatus === 2) {
                html += '<div style="padding:16px;text-align:center;color:var(--layer8d-text-medium);">';
                html += evalImport.processingStatus === 1 ? 'Waiting to process...' : 'Extracting findings from document...';
                html += '</div>';
                html += '</div>';
                return html;
            }

            if (findings.length === 0) {
                html += '<div style="padding:16px;text-align:center;color:var(--layer8d-text-medium);">No findings extracted.</div>';
            } else {
                html += '<h4 style="margin:12px 0 8px;">Findings (' + findings.length + ')</h4>';
                html += '<table class="l8-eval-findings-table" style="width:100%;border-collapse:collapse;font-size:13px;">';
                html += '<thead><tr>';
                html += '<th style="text-align:left;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">Section</th>';
                html += '<th style="text-align:left;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">Field</th>';
                html += '<th style="text-align:left;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">New Value</th>';
                html += '<th style="text-align:center;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">Confidence</th>';
                html += '<th style="text-align:center;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">Status</th>';
                html += '<th style="text-align:center;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">Actions</th>';
                html += '</tr></thead><tbody>';

                for (var i = 0; i < findings.length; i++) {
                    html += this.renderFindingRow(findings[i], i);
                }
                html += '</tbody></table>';
            }

            if (contradictions.length > 0) {
                html += '<h4 style="margin:12px 0 8px;color:var(--layer8d-warning);">Contradictions (' + contradictions.length + ')</h4>';
                for (var c = 0; c < contradictions.length; c++) {
                    var ct = contradictions[c];
                    html += '<div style="padding:8px;margin:4px 0;border:1px solid var(--layer8d-warning);border-radius:4px;">';
                    html += '<strong>' + Layer8DUtils.escapeHtml(ct.profileField) + '</strong><br>';
                    html += 'Current: ' + Layer8DUtils.escapeHtml(ct.currentValue || '-') + '<br>';
                    html += 'Document says: ' + Layer8DUtils.escapeHtml(ct.documentSays || '-') + '<br>';
                    html += 'AI recommends: ' + Layer8DUtils.escapeHtml(ct.aiRecommendation || '-');
                    html += '</div>';
                }
            }

            html += '</div>';
            return html;
        },

        renderFindingRow: function(finding, index) {
            var esc = Layer8DUtils.escapeHtml;
            var conf = finding.confidence || 0;
            var confClass = conf >= 0.8 ? 'layer8d-status-active' : (conf >= 0.5 ? 'layer8d-status-pending' : 'layer8d-status-inactive');
            var confPct = Math.round(conf * 100) + '%';
            var statusLabel = StudentsPeople.enums.EVAL_FINDING_STATUS[finding.status] || 'Pending';

            var html = '<tr data-finding-index="' + index + '">';
            html += '<td style="padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">' + esc(finding.profileSection) + '</td>';
            html += '<td style="padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">' + esc(finding.profileField) + '</td>';
            html += '<td style="padding:4px 8px;border-bottom:1px solid var(--layer8d-border);" title="' + esc(finding.sourceText || '') + '">' + esc(finding.newValue || '-') + '</td>';
            html += '<td style="text-align:center;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);"><span class="' + confClass + '" style="padding:2px 6px;border-radius:3px;">' + confPct + '</span></td>';
            html += '<td style="text-align:center;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">' + esc(statusLabel) + '</td>';
            html += '<td style="text-align:center;padding:4px 8px;border-bottom:1px solid var(--layer8d-border);">';
            html += '<button class="l8-eval-accept-btn layer8d-btn layer8d-btn-small" data-idx="' + index + '" data-action="accept" style="margin:1px;">Accept</button>';
            html += '<button class="l8-eval-reject-btn layer8d-btn layer8d-btn-small" data-idx="' + index + '" data-action="reject" style="margin:1px;">Reject</button>';
            html += '</td>';
            html += '</tr>';
            return html;
        },

        onFindingAction: function(evalImport, findingIndex, action) {
            var finding = evalImport.findings[findingIndex];
            if (!finding) return;
            if (action === 'accept') {
                finding.status = 2; // ACCEPTED
            } else if (action === 'reject') {
                finding.status = 3; // REJECTED
            }
            this.recomputeCounts(evalImport);
        },

        recomputeCounts: function(evalImport) {
            var accepted = 0, rejected = 0, allReviewed = true;
            var findings = evalImport.findings || [];
            for (var i = 0; i < findings.length; i++) {
                if (findings[i].status === 2) accepted++;
                else if (findings[i].status === 3) rejected++;
                else if (findings[i].status === 1) allReviewed = false;
            }
            evalImport.acceptedCount = accepted;
            evalImport.rejectedCount = rejected;
            evalImport.allReviewed = allReviewed;
        },

        buildApplyPayload: function(evalImport) {
            evalImport.appliedToProfile = true;
            return evalImport;
        }
    };
})();
