USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaIRType]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaIRType]
AS
SELECT        Code, Style, AllCurrencies, CurrencyIsoCode, RateCategory
FROM            {{!SQL.SOURCE}}.IRType
WHERE        (InternalDeleted IS NULL)
GO
