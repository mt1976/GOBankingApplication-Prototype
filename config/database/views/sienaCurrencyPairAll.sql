USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCurrencyPairAll]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCurrencyPairAll]
AS
SELECT        {{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode, {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode, {{!SQL.SOURCE}}.CurrencyPair.Active, {{!SQL.SOURCE}}.CurrencyPair.ReciprocalActive, 
                         {{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode + {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode AS Code, {{!SQL.SOURCE}}.Currency.Name AS MajorName, Currency_1.Name AS MinorName
FROM            {{!SQL.SOURCE}}.CurrencyPair INNER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode = {{!SQL.SOURCE}}.Currency.Code INNER JOIN
                         {{!SQL.SOURCE}}.Currency AS Currency_1 ON {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode = Currency_1.Code
WHERE        ({{!SQL.SOURCE}}.CurrencyPair.InternalDeleted IS NULL)
GO
