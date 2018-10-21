<?php

class News_model extends CI_Model
{
	public $table;
	
	public function __construct()
	{
		$this->load->database();
		
		$this->table = 'news';
	}
	
	public function edit_news()
	{
		$this->load->helper('url');

		$alias = url_title($this->input->post('alias'), 'dash', TRUE);
		$content = $this->input->post('content');
		$id = $this->input->post('id');
		$title = $this->input->post('title');

		$data = array
		(
			'alias' => $alias,
			'title' => $title,
			'content' => $content
		);
		
		$this->db->where('id', $id);
		$result = $this->db->update($this->table, $data);

		return $result;
	}
	
	public function get_news_all()
	{
		$query = $this->db->get($this->table);
		$result = $query->result_array();
		
		return $result;
	}
	
	public function get_news_titles_all()
	{
		$this->db->select('id, alias, title');
		$query = $this->db->get($this->table);
		$result = $query->result_array();
		
		return $result;
	}
	
	public function get_news_by_alias($alias)
	{
		$query_param = array('alias' => $alias);
		$query = $this->db->get_where($this->table, $query_param);
		$result = $query->row_array();
		
		return $result;
	}
	
	public function get_news_by_id($id)
	{
		$query_param = array('id' => $id);
		$query = $this->db->get_where($this->table, $query_param);
		$result = $query->row_array();
		
		return $result;
	}
	
	public function insert_news()
	{
		$this->load->helper('url');

		$alias = url_title($this->input->post('title'), 'dash', TRUE);
		$content = $this->input->post('content');
		$title = $this->input->post('title');

		$data = array
		(
			'alias' => $alias,
			'title' => $title,
			'content' => $content
		);
		
		$result = $this->db->insert($this->table, $data);

		return $result;
	}
}
