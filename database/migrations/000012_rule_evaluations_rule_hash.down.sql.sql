-- Copyright 2024 Stacklok, Inc
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- Begin transaction
BEGIN;

-- Prevent concurrent updates to rule_evaluations
SELECT *
FROM rule_evaluations FOR UPDATE;

-- Delete duplicate rule evaluation results without considering rule_hash
-- Using CTID as postgres doesn't have min, max aggregators for uuid (too much code to add one)
DELETE
FROM rule_evaluations
WHERE CTID IN (SELECT MIN(CTID) AS CTID
               FROM rule_evaluations
               GROUP BY entity, profile_id, repository_id, rule_type_id,
                        COALESCE(pull_request_id, '00000000-0000-0000-0000-000000000000'::UUID),
                        COALESCE(artifact_id, '00000000-0000-0000-0000-000000000000'::UUID)
               HAVING COUNT(*) > 1);


-- Drop the existing unique index on rule_evaluations
DROP INDEX IF EXISTS rule_evaluations_results_idx;

-- Recreate the unique index without rule_hash
CREATE UNIQUE INDEX rule_evaluations_results_idx
    ON rule_evaluations (profile_id, repository_id, COALESCE(artifact_id, '00000000-0000-0000-0000-000000000000'::UUID),
                         entity, rule_type_id, COALESCE(pull_request_id, '00000000-0000-0000-0000-000000000000'::UUID));

-- Remove the rule_hash column from rule_evaluations
ALTER TABLE rule_evaluations
    DROP COLUMN IF EXISTS rule_hash;

-- Commit transaction
COMMIT;
