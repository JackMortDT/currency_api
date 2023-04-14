CREATEDATABASE currency_devCREATETABLEpublic.call_records(
    idBIGINTNOTNULL,
    created_atTIMESTAMPWITHTIMEZONE,
    updated_atTIMESTAMPWITHTIMEZONE,
    deleted_atTIMESTAMPWITHTIMEZONE,
    request_dateTIMESTAMPWITHTIMEZONE,
    durationBIGINT,
    sucessBOOLEAN,
    successBOOLEAN
);CREATESEQUENCEpublic.call_records_id_seqSTARTWITH1INCREMENTBY1NOMINVALUENOMAXVALUE CACHE1;ALTERSEQUENCEpublic.call_records_id_seq OWNEDBYpublic.call_records.id;CREATETABLEpublic.currency_rates(
    idBIGINTNOTNULL,
    codeTEXT,
    VALUENUMERIC,
    created_atTIMESTAMPWITHTIMEZONE,
    updated_atTIMESTAMPWITHTIMEZONE
);CREATESEQUENCEpublic.currency_rates_id_seqSTARTWITH1INCREMENTBY1NOMINVALUENOMAXVALUE CACHE1;ALTERSEQUENCEpublic.currency_rates_id_seq OWNEDBYpublic.currency_rates.id;ALTERTABLEONLYpublic.call_recordsALTERCOLUMNid
SETDEFAULTnextval('public.call_records_id_seq'::REGCLASS);ALTERTABLEONLYpublic.currency_ratesALTERCOLUMNid
SETDEFAULTnextval('public.currency_rates_id_seq'::REGCLASS);CREATEINDEX idx_call_records_deleted_at
ON public.call_records USING btree(
deleted_at
);CREATEUNIQUEINDEX idx_code_updated_at
ON public.currency_rates USING btree(
    code,
    updated_at
);
