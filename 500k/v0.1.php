<?php

/*
	Simple Electronic Mail Address Checker.
	
	Reads E-mail Addresses from a Table of the MySQL DataBase, 
	For each Address checks:
	
		- its Syntax, 
		- Accessibility of Domain, 
		- Existance of the Username at Mail Server.
	
	List of Functions which are not implemented:
	
		1.	username_clean 
		
				Reason:		depends on many RFCs and Restrictions of 
							each particular Mail Server. Can not be done in 30 Minutes as suggested.
		
		2.	username_check_syntax
		
				Reason:		depends on many RFCs and Restrictions of 
							each particular Mail Server. Can not be done in 30 Minutes as suggested.
		
		3.	address_check_access
		
				Reason:		Can not be done in 30 Minutes as suggested.
	
	Version:	0.1.
	Date:		2018-01-14.
	Author:		McArcher.
*/

// Constants.
const CRLF = "\r\n";
const NL = "<br>\r\n";

// Parameters.
$debug = true;
$rfc_1123_isEnabled = false;

// DataBase.
$db; // DB Handler.
$db_base	= 'test_1';
$db_charset	= 'utf8';
//$db_host	= 'localhost';
$db_host	= '127.0.0.1';
$db_port	= 3306;
$db_pwd		= 'test_1';
$db_user	= 'test_1';

// Variables.
$addr_clear = null;
$addr_raw = null;
$checker = null;
$i = null;

//-----------------------------------------------------------------------------|

class MailAddressChecker
{
	private $base = null;				// DataBase Name.
	private $db = null;					// DataBase Handle.
	private $debug = null;				// Debug Mode (Flag).
	private $rfc_1123_isEnabled = null; // RFC 1123 Enabled (Flag).
	
	private const COL_ID = 'i_id';		// Column : ID.
	private const COL_ADDR = 'm_mail';	// Column : E-Mail Address.
	private const TABLE = 'emails';		// Table in D.B.
	private const HOSTNAME_LEN_MAX = 253;	// Maximum Length of Domains String.
	private const DOMAIN_LEN_MAX = 63;	// Maximum Length of Sub-Domain.
	private const DNS_RECORD_TYPE_MX = 'MX';
	
	public $addr_valid = array();		// Valid Flag for each Address.
	public $addresses_clear = null; 	// Addresses without Comments.
	public $addresses_raw = null; 		// Addresses read from DataBase.
	
	// List of inaccessible Domains with correct Syntax.
	public $domains_bad = array();
	
	// List of accessible Domains with correct Syntax.
	public $domains_good = array();
	
	public $msg = '';					// Last Action's Report.
	
	// Constructor.
	public function __construct
	(
		$db_host, 
		$db_user, 
		$db_pwd, 
		$db_base, 
		$db_port = 3306, 
		$db_charset = 'utf8', 
		$debug = false,
		$rfc_1123_isEnabled = false
	)
	{
		$this->debug = $debug;
		$this->base = $db_base;
		$this->rfc_1123_isEnabled = $rfc_1123_isEnabled;
		
		$connect_result = null;
		$disconnect_result = null;
		$read_result = null;
		
		// 1. Connect to DataBase.
		$connect_result =
			$this->db_connect
			(
				$db_host, 
				$db_user, 
				$db_pwd, 
				$db_base, 
				$db_port, 
				$db_charset
			);
		
		if ($debug)
		{
			echo $this->msg . NL;
		}
		
		if ($connect_result === false)
		{
			return false;
		}
		
		// 2. Read all Records.
		$read_result = $this->db_read();
		
		if ($debug)
		{
			echo $this->msg . NL;
		}
		
		// 3. DisConnect from DataBase.
		$disconnect_result = $this->db_disconnect();
		
		if ($debug)
		{
			echo $this->msg . NL;
		}
		
		if ($disconnect_result === false)
		{
			return false;
		}
		
		// 4. Process Addresses.
		$this->addresses_process();
	}
	
	// Connect to DataBase.
	private function db_connect
	(
		$db_host, 
		$db_user, 
		$db_pwd, 
		$db_base, 
		$db_port, 
		$db_charset
	)
	{
		$db = null;
		$debug = null;
		
		$debug = $this->debug;
		
		$db = new mysqli($db_host, $db_user, $db_pwd, $db_base, $db_port);
		
		if ($db->connect_errno)
		{
			if ($debug)
			{
				$this->msg =
					'Connection to "' . $db_host . '" has failed. ' . 
					'Error (' . $db->connect_errno . '): ' . 
					$db->connect_error;
			}
			
			$this->db = null;
			
			return false;
		}
		
		$db->set_charset($db_charset);
		
		if ($debug)
		{
			$this->msg = 'Connected to ' . $db->host_info . '.';
		}
		
		$this->db = $db;
		
		return true;
	}
	
	// DisConnect from DataBase.
	private function db_disconnect()
	{
		$debug = null;
		$result = null;
		
		$debug = $this->debug;
		
		$result = $this->db->close();
		
		if ($result)
		{
			if ($debug)
			{
				$this->msg = 'DisConnect was successful.';
			}
			
			return true;
		}
		else
		{
			if ($debug)
			{
				$this->msg = 'DisConnect has failed.';
			}
			
			return false;
		}
	}
	
	// Read all Addresses from DataBase.
	private function db_read()
	{
		$addresses_count = null;
		$cmd = null;				// Current SQL Query.
		$cmd_result = null;
		$debug = null;
		$row = null;
		$table_full = null;
		
		$debug = $this->debug;
		$table_full = $this->base . '.' . self::TABLE;
		
		// Get Number of all Addresses in DataBase.
		$cmd = 'SELECT COUNT(*) FROM ' . $table_full . ';';
		$cmd_result = $this->db->query($cmd);
		if ($cmd_result === false)
		{
			if ($debug)
			{
				$this->msg =	'Error during DataBase reading: ' . 
								$this->db->error . '.';
			}
			
			return false;
		}
		
		$row = $cmd_result->fetch_array();
		$addresses_count = $row[0];
		$cmd_result->close();
		
		// Get all the Records at once.
		$cmd =	'SELECT ' . self::COL_ID . ', ' . self::COL_ADDR . ' ' . 
				'FROM ' . $table_full . ';';
		$cmd_result = $this->db->query($cmd, MYSQLI_STORE_RESULT);
		if ($cmd_result === false)
		{
			if ($debug)
			{
				$this->msg =	'Error during DataBase reading: ' . 
								$this->db->error . '.';
			}
			
			return false;
		}
		$this->addresses_raw = $cmd_result->fetch_all();
		$cmd_result->close();
		
		if ($debug)
		{
			$this->msg =	$addresses_count . ' Records were found.';
		}
		
		return true;
	}
	
	// Process all Addresses.
	private function addresses_process()
	{
		$addr = null;
		$addr_clear = null;
		$addr_isAccessible = null;
		$at_pos_last = null;
		$debug = null;
		$domain_clear = null;
		$domain_isAccessible = null;
		$domain_isCorrect = null;
		$domain_raw = null;
		$domain_idx_good = null;
		$domain_idx_bad = null;
		$end = null;
		$i = null;
		$row = null;
		$start = null;
		$tmp = null;
		$username_clear = null;
		$username_raw = null;
		$username_syntax_isGood = null;
		
		$debug = $this->debug;
		
		foreach ($this->addresses_raw as $i => $row)
		{
			$addr = $row[1];
			$at_pos_last = strrpos($addr, '@');
			
			if ($debug)
			{
				echo	NL . 
						$i . '. ' . $addr . '.' . NL;//
			}
			
			// 1. Check Syntax.
			// 1.1. Last '@' not found ?
			if ($at_pos_last === false)
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			$username_raw = substr($addr, 0, $at_pos_last);
			$start = $at_pos_last + 1;
			$end = strlen($addr) - $start;
			$domain_raw = substr($addr, $start, $end);
			
			// 1.2. Empty Username or Domain ?
			if ((strlen($username_raw) == 0) || (strlen($domain_raw) == 0))
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			// 1.3. Remove Comments from Domain.
			$domain_clear = $this->domain_clean($domain_raw);
			if ($domain_clear === false)
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			// 1.4. Check Characters of Domain.
			$domain_isCorrect = $this->domain_check_syntax($domain_clear);
			if ($domain_isCorrect === false)
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			if ($debug)
			{
				echo	$domain_clear . '.' . NL . 
						'Synatax is Good.' . NL;//
			}
			
			// 2. Have already seen this Domain ?
			$domain_idx_good =
				array_search($domain_clear, $this->domains_good);
			
			if ($domain_idx_good === false)
			{
				// Domain is not in Good List.
				$domain_idx_bad =
					array_search($domain_clear, $this->domains_bad);
				
				if ($domain_idx_bad === false)
				{
					// Domain is new (not seen before).
					
					if ($debug)
					{
						echo 'Domain is new.' . NL;//
					}
					
					// Check Accessibility of Domain.
					$domain_isAccessible = 
						$this->domain_check_access($domain_clear);
					
					if ($domain_isAccessible === false)
					{
						$this->addr_valid[$i] = false;
						
						// Add Domain to Bad Domains List.
						$this->domains_bad[] = $domain_clear;
						
						continue;
					}
					
					// Domain is Accessible.
					// Add Domain to Good Domains List.
					$this->domains_good[] = $domain_clear;
					
					if ($debug)
					{
						echo 'Domain is accessible.' . NL;//
					}
				}
				else
				{
					// Domain is in Bad List.
					
					$this->addr_valid[$i] = false;
					
					if ($debug)
					{
						echo 'Domain is already marked as Bad.' . NL;//
					}
				
					continue;
				}
			}
			else
			{
				// Domain is in Good Domains List.
				
				if ($debug)
				{
					echo 'Domain is already marked as Good.' . NL;//
				}
			}
			
			// Domain is Good.
			// Now, check the Username.
			
			// Remove Comments from Username.
			$username_clear = $this->username_clean($username_raw);
			if ($username_clear === false)
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			// Check Syntax of the Username.
			$username_syntax_isGood =
				$this->username_check_syntax($username_clear);
			
			if ($username_syntax_isGood === false)
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			// Save Clear Address into Object.
			$addr_clear = $username_clear . '@' . $domain_clear;
			$this->addresses_clear[$i] = $addr_clear;
			
			// Check Accessibility of E-Mail Address.
			$addr_isAccessible =
				$this->address_check_access($addr_clear);
			
			if ($addr_isAccessible === false)
			{
				$this->addr_valid[$i] = false;
				
				continue;
			}
			
			$this->addr_valid[$i] = true;
		}
	}
	
	// Removes all Comments from Domain.
	private function domain_clean($domain_raw)
	{
		$domain_raw_len = null;
		$domain_tmp = null;
		$domain_tmp_len = null;
		$bkt_left_first_pos = null;		// Position of first '('.
		$bkt_right_first_pos = -1;		// Position of first ')'.
		$bkt_left_last_pos = null;		// Position of last '('.
		$bkt_right_last_pos = null;		// Position of last ')'.
		$comment_left_exists = false;
		$comment_right_exists = false;
		
		$debug = $this->debug;
		$domain_raw_len = strlen($domain_raw);
		
		// 1. Left Corner.
		$bkt_left_first_pos = strpos($domain_raw, '(');
		if	(
				($bkt_left_first_pos !== false) &&
				($bkt_left_first_pos === 0)
			)
		{
			$comment_left_exists = true;
			$bkt_right_first_pos = strpos($domain_raw, ')');
			
			if ($bkt_right_first_pos === false)
			{
				if ($debug)
				{
					$this->msg =	'Syntax Error in Domain: ' . 
									'Closing Bracket is not found.';
				}
				
				return false;
			}
			else
			{
				// Closing Bracket is found.
				// Cut left Bracket.
				$domain_tmp = substr($domain_raw, 0, $bkt_right_first_pos + 1);
			}
		}
		else
		{
			$domain_tmp = $domain_raw;
		}
		
		$domain_tmp_len = strlen($domain_tmp);
		
		// 2. Right Corner.
		$bkt_right_last_pos = strrpos($domain_tmp, ')');
		
		if	(
				($bkt_right_last_pos !== false) &&
				($bkt_right_last_pos === $domain_tmp_len - 1)
			)
		{
			$comment_right_exists = true;
			$bkt_left_last_pos = strrpos($domain_tmp, '(');
			
			if ($bkt_left_last_pos === false)
			{
				if ($debug)
				{
					$this->msg =	'Syntax Error in Domain: ' . 
									'Opening Bracket is not found.';
				}
				
				return false;
			}
			else
			{
				// Opening Bracket is found.
				// Cut right Bracket.
				$domain_tmp = substr($domain_tmp, 0, $bkt_left_last_pos);
			}
		}
		
		// 3. Center Part.
		$bkt_left_first_pos = strpos($domain_tmp, '(');
		$bkt_right_last_pos = strrpos($domain_tmp, ')');
		
		if	(
				($bkt_left_first_pos !== false) ||
				($bkt_right_last_pos !== false)
			)
		{
			// Bracket in the Middle!
			
			if ($debug)
			{
				$this->msg =	'Syntax Error in Domain: ' . 
								'Bracket in the Middle of String.';
			}
			
			return false;
		}
		
		return $domain_tmp;
	}
	
	// Checks the Syntax of Domain.
	private function domain_check_syntax($address)
	{
		// /!\ Internationalized Domain Names are ignored /!\
		
		$address_len = null;
		$domain = null;
		$domain_len = null;
		$domain_idx = null;
		$domain_last_idx = null;
		$domains = null;
		$domains_count = null;
		$regexp_pattern = null;
		$regexp_result = null;
		
		$debug = $this->debug;
		$rfc_1123_isEnabled = $this->rfc_1123_isEnabled;
		
		// Domains String Length.
		$address_len = strlen($address);
		if ($address_len > self::HOSTNAME_LEN_MAX)
		{
			if ($debug)
			{
				$this->msg =	'Error: Domain is longer than ' . 
								self::HOSTNAME_LEN_MAX . 
								' Symbols.';
			}
			
			return false;
		}
		
		$domains = explode('.', $address);
		$domains_count = sizeof($domains);
		$domain_last_idx = $domains_count - 1;
		
		foreach ($domains as $domain_idx => $domain)
		{
			$domain_len = strlen($domain);
			
			// Empty not-last Domain ?
			if	(
					($domain_len == 0) &&
					($domain_idx != $domain_last_idx)
				)
			{
				if ($debug)
				{
					$this->msg =	'Syntax Error in Domain: ' . 
									'Empty non-last Domain.';
				}
				
				return false;
			}
			
			$domain = strtolower($domain);
			
			// Check Characters.
			if ($rfc_1123_isEnabled)
			{
				// See RFC 1123.
				$regexp_pattern =
					'/^' . 
					'[0-9\-a-zA-Z]{1,' . self::DOMAIN_LEN_MAX . '}' . 
					'$/';
			}
			else
			{
				// See RFC 952.
				$regexp_pattern =
					'/^' . 
					'[a-zA-Z]{1,1}' . 
					'[0-9\-a-zA-Z]{0,' . (self::DOMAIN_LEN_MAX - 2) . '}' . 
					'[0-9a-zA-Z]{1,1}' . 
					'$/';
			}
			if ($domain_len > 0)
			{
				$regexp_result = preg_match($regexp_pattern, $domain);
				
				if ($regexp_result !== 1)
				{
					if ($debug)
					{
						$this->msg =	'Syntax Error in Domain "' . 
										$address . '": Bad Sub-Domain: ' . 
										$domain;
					}
					
					return false;
				}
			}
		}
		
		return true;
	}
	
	// Checks the Accessibility of Domain.
	private function domain_check_access($address)
	{
		$i = null;
		$record = null;
		$type_cur = null;
		
		$debug = $this->debug;
		$reply = @dns_get_record($address, DNS_MX);
		
		// Fail ?
		if ($reply === false)
		{
			if ($debug)
			{
				$this->msg =	'Failed to access Domain Name "' . 
								$address . '".';
			}
			
			return false;
		}
		
		// Read Reply.
		foreach ($reply as $i => $record)
		{
			$type_cur = $record['type'];
			
			if ($type_cur == self::DNS_RECORD_TYPE_MX)
			{
				// MX is Found!
				
				return true;
			}
		}
		
		if ($debug)
		{
			$this->msg =	'Domain Name "' . $address . 
							'" has no MX Records in DNS.';
		}
		
		return false;
	}
	
	// Removes all Comments from the Username.
	private function username_clean($username_raw)
	{
		$username_clear = null;
		
		///
		$username_clear = $username_raw; // Затычка.
		///
		
		return $username_clear;
	}
	
	// Checks Syntax of the Username.
	private function username_check_syntax($username)
	{
		///
		return true; // Затычка.
	}
	
	// Checks Accessibility of the E-Mail Address.
	private function address_check_access($address)
	{
		///
		return true; // Затычка.
	}
}

//-----------------------------------------------------------------------------|

// Do the Job.
$checker = 
	new MailAddressChecker(	$db_host, $db_user, $db_pwd, $db_base, $db_port, 
							$db_charset, $debug, $rfc_1123_isEnabled);

// Show Results.
echo NL . '[Results]' . NL . NL;

echo 'RFC 1123 is ';
if ($rfc_1123_isEnabled)
{
	echo 'enabled.' . NL . NL;
}
else
{
	echo 'disabled.' . NL . NL;
}

foreach ($checker->addresses_raw as $i => $addr_raw)
{
	if ($checker->addr_valid[$i])
	{
		$addr_clear = $checker->addresses_clear[$i];
		
		echo	$i . '. ' . $addr_raw[1] . ' ( ' . $addr_clear . ' ) : ' . 
				'Good.' . NL;
	}
	else
	{
		echo	$i . '. ' . $addr_raw[1] . ' : Bad.' . NL;
	}
}

?>
