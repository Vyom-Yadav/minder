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

-- Add rule_hash to rule_evaluations
ALTER TABLE rule_evaluations
    ADD COLUMN rule_hash TEXT;

-- Drop the existing unique index on rule_evaluations
DROP INDEX IF EXISTS rule_evaluations_results_idx;

-- Recreate the unique index with rule_hash
CREATE UNIQUE INDEX rule_evaluations_results_idx
    ON rule_evaluations (profile_id, repository_id, COALESCE(artifact_id, '00000000-0000-0000-0000-000000000000'::UUID),
                         entity, rule_type_id, COALESCE(pull_request_id, '00000000-0000-0000-0000-000000000000'::UUID),
                         rule_hash);

CREATE OR REPLACE FUNCTION compute_rule_hash(entity entities, profile_id uuid, rule_type_id uuid) RETURNS text AS
$$
DECLARE
    rule_hash text;
    rule_name text;
    rules     jsonb;
BEGIN
    SELECT entity_profiles.contextual_rules
    INTO rules
    FROM entity_profiles
    WHERE entity_profiles.profile_id = compute_rule_hash.profile_id
      AND entity_profiles.entity = compute_rule_hash.entity;

    SELECT rule_type.name INTO rule_name FROM rule_type WHERE id = compute_rule_hash.rule_type_id;

    SELECT encode(sha256(to_jsonb(rule_element) :: TEXT :: BYTEA), 'hex')
    INTO rule_hash
    FROM jsonb_array_elements(rules) rule_element
    WHERE rule_element ->> 'type' = rule_name;
    RETURN rule_hash;
END;
$$ LANGUAGE plpgsql STABLE;

-- Prevent rule_evaluations to be updated outside the transaction
SELECT *
FROM rule_evaluations FOR UPDATE;

-- Prevent entity_profiles to be updated outside the transaction
SELECT *
FROM entity_profiles FOR UPDATE;

-- Update rule evaluations
UPDATE rule_evaluations
SET rule_hash = compute_rule_hash(rule_evaluations.entity, rule_evaluations.profile_id, rule_evaluations.rule_type_id)
WHERE true;

-- Add non null constraint on rule_hash
ALTER TABLE rule_evaluations
    ALTER COLUMN rule_hash SET NOT NULL;

-- transaction commit
COMMIT;
