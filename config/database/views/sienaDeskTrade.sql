USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDeskTrade]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDeskTrade]
AS
SELECT        Name
FROM            {{!SQL.SOURCE}}.TraderTradingEntity
WHERE        (InternalDeleted IS NULL)
GO
