USE [SRS]
GO
/****** Object:  Table [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[sienadbStore]    Script Date: 10/05/2021 22:16:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[sienadbStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[ID] [nvarchar](max) NOT NULL,
	[server] [nvarchar](max) NULL,
	[port] [nvarchar](max) NULL,
	[user] [nvarchar](max) NULL,
	[password] [nvarchar](max) NULL,
	[database] [nvarchar](max) NULL,
	[schema] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[active] [nvarchar](max) NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
