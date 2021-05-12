USE [SRS]
GO
/****** Object:  Table [dbo].[rateDataStore]    Script Date: 10/05/2021 22:16:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[rateDataStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[id] [nvarchar](max) NOT NULL,
	[bid] [nvarchar](max) NULL,
	[mid] [nvarchar](max) NULL,
	[offer] [nvarchar](max) NULL,
	[market] [nvarchar](max) NULL,
	[tenor] [nvarchar](max) NULL,
	[series] [nvarchar](max) NULL,
	[name] [nvarchar](max) NULL,
	[source] [nvarchar](max) NULL,
	[destination] [nvarchar](max) NULL,
	[class] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[date] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
 CONSTRAINT [PK_rateDataStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO