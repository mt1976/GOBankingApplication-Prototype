USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyGroup]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyGroup]
AS
SELECT        {{!SQL.SOURCE}}.CpartyGroup.Name, {{!SQL.SOURCE}}.CpartyGroup.CountryCode, {{!SQL.SOURCE}}.CpartyGroup.SuperGroup
FROM            {{!SQL.SOURCE}}.CpartyGroup 
WHERE        ({{!SQL.SOURCE}}.CpartyGroup.InternalDeleted IS NULL)
GO
