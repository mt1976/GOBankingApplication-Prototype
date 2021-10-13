USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealType]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealType]
AS
SELECT        {{!SQL.SOURCE}}.DealType.DealTypeKey, {{!SQL.SOURCE}}.DealType.DealTypeShortName, {{!SQL.SOURCE}}.DealType.FundamentalDealTypeKey, {{!SQL.SOURCE}}.DealType.RFQ, {{!SQL.SOURCE}}.DealType.RelatedDealType, {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.ProductCode
FROM            {{!SQL.SOURCE}}.DealType LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.FundamentalDealTypeSwitches ON {{!SQL.SOURCE}}.DealType.DealTypeKey = {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.DealTypeKey
WHERE        ({{!SQL.SOURCE}}.DealType.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealType.IsActive = 1) AND ({{!SQL.SOURCE}}.DealType.Interbook = 0)
GO
