USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[contactManagerDealRefNo]    Script Date: 24/11/2021 20:38:01 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[contactManagerDealRefNo]
AS
SELECT        refNo, noteId, recordState
FROM            CM.DealRefNo
GO