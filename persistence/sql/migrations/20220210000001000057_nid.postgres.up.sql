-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrate gen

CREATE UNIQUE INDEX hydra_jwk_sid_kid_nid_key ON hydra_jwk (sid ASC, kid ASC, nid ASC);

