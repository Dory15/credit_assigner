
USE [credit_assigner]
GO
/****** Object:  StoredProcedure [dbo].[insert_credit_assigment_statistics]    Script Date: 13/03/2024 04:16:04 p. m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		<Dorian Fernandez>
-- Create date: <13/03/2024>
-- Description:	<SP to insert credit assigment statistics>
-- =============================================
ALTER PROCEDURE [dbo].[insert_credit_assigment_statistics]
(
	@assigment_succesful BIT 
)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
	SET TRANSACTION ISOLATION LEVEL READ COMMITTED 

	BEGIN TRY 
	
		 DECLARE @trancount INT;
		 SET @trancount = @@TRANCOUNT;
			IF @trancount = 0
				BEGIN TRAN [tran_insert_statistics]; 
			ELSE
				SAVE TRANSACTION [tran_insert_statistics];

		DECLARE	@assignments_made INT, @successful_assignments INT, @failed_assignments INT, @average_successful_assignments FLOAT, @average_failed_assignments FLOAT;

		IF NOT EXISTS(SELECT 1 FROM credit_statistics)
		BEGIN
			INSERT INTO credit_statistics (assignments_made, successful_assignments, failed_assignments, average_successful_assignments, average_failed_assignments) VALUES (1, @assigment_succesful, 1-@assigment_succesful, @assigment_succesful*100, (1-@assigment_succesful)*100);
		END
		ELSE
		BEGIN

		SELECT 
		@assignments_made = assignments_made,
		@successful_assignments = successful_assignments,
		@failed_assignments = failed_assignments,
		@average_successful_assignments = average_successful_assignments,
		@average_failed_assignments = average_failed_assignments
		FROM credit_statistics;

		SET @assignments_made = @assignments_made + 1;
		SET @successful_assignments = @successful_assignments + @assigment_succesful;
		SET @failed_assignments = @failed_assignments + (1 - @assigment_succesful);
		SET @average_successful_assignments = ( @successful_assignments  * 100.0 )/@assignments_made;
		SET @average_failed_assignments = (@failed_assignments * 100.0)/@assignments_made ;

		UPDATE credit_statistics SET assignments_made = @assignments_made, successful_assignments = @successful_assignments, failed_assignments = @failed_assignments, average_successful_assignments = @average_successful_assignments, average_failed_assignments = @average_failed_assignments WHERE id = 1;

		END

		IF @trancount = 0
			COMMIT TRAN [tran_insert_statistics];

	END TRY
	BEGIN CATCH
		DECLARE @xstate INT

		SELECT @xstate = XACT_STATE();

		IF @xstate = -1
			ROLLBACK;
		IF @xstate = 1
			ROLLBACK TRANSACTION [tran_insert_statistics];
	END CATCH
	
		SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED 
END
