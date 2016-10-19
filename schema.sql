CREATE TABLE "pcs-application-tbl" (
    _id text NOT NULL,
    profileid text,
    name text,
    answers json,
    status text,
    notification json,
    createddate timestamp with time zone DEFAULT now(),
    modifieddate timestamp with time zone,
    modifiedby text
);

CREATE TABLE "pcs-document-tbl"
(
  _id text NOT NULL,
  label text,
  uploadid text NOT NULL,
  filename text NOT NULL,
  createddate timestamp with time zone NOT NULL DEFAULT now(),
  createdby text NOT NULL,
  contenttype text NOT NULL,
  CONSTRAINT "pcs-document-tbl_pkey" PRIMARY KEY (_id)
);

CREATE TABLE "pcs-profile"
(
  _id text NOT NULL,
  name text,
  email text,
  sfdcid text,
  CONSTRAINT "pcs-profile_pkey" PRIMARY KEY (_id)
);

CREATE TABLE "pcs-question-tbl"
(
  _id text NOT NULL,
  title text NOT NULL,
  name text,
  description text NOT NULL,
  type integer NOT NULL DEFAULT 1001,
  "answerOptions" json NOT NULL DEFAULT '{}'::json,
  priority integer NOT NULL DEFAULT (-1),
  status text NOT NULL DEFAULT 'pending'::text,
  CONSTRAINT "pcs-question-tbl_pkey" PRIMARY KEY (_id)
);
