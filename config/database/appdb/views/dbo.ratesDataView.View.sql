USE [SRS-2]
GO
/****** Object:  View [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[ratesDataView] AS select * from [{{!SQL.SOURCE}}].dbo.rateDataStore rds

GO
