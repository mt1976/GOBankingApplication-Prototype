USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaImportOrigins]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaImportOrigins]
AS
SELECT DISTINCT KeyOriginID
FROM            {{!SQL.SOURCE}}.CounterpartyImportID
GO
