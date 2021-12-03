USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealQuoteHistory]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealQuoteHistory]
AS
SELECT        TOP (100) PERCENT {{!SQL.SOURCE}}.QuoteHistory.SienaReference, {{!SQL.SOURCE}}.QuoteHistory.QuoteNo, {{!SQL.SOURCE}}.QuoteHistory.ContractNumber, {{!SQL.SOURCE}}.QuoteHistory.QuoteTimestamp, {{!SQL.SOURCE}}.QuoteHistory.UTCQuoteTimestamp, {{!SQL.SOURCE}}.QuoteHistory.EventType, 
                         {{!SQL.SOURCE}}.QuoteHistoryRate.RateNo, {{!SQL.SOURCE}}.QuoteHistoryRate.Category, {{!SQL.SOURCE}}.QuoteHistoryRate.RateType, {{!SQL.SOURCE}}.QuoteHistoryRate.RateKey, {{!SQL.SOURCE}}.QuoteHistoryRate.RateValue, {{!SQL.SOURCE}}.QuoteHistoryRate.RateSource
FROM            {{!SQL.SOURCE}}.QuoteHistory LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.QuoteHistoryRate ON {{!SQL.SOURCE}}.QuoteHistory.QuoteHistoryId = {{!SQL.SOURCE}}.QuoteHistoryRate.QuoteHistoryRateId
ORDER BY {{!SQL.SOURCE}}.QuoteHistory.SienaReference, {{!SQL.SOURCE}}.QuoteHistory.QuoteNo
GO
