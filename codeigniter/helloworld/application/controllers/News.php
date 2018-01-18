<?php
defined('BASEPATH') OR exit('No direct script access allowed');

class News extends CI_Controller
{
	public $form_add_submit_alias;
	public $form_edit_submit_alias;
	
	public function __construct()
	{
		parent::__construct();
		
		$this->load->model('news_model');
		$this->load->helper('url_helper');
		
		$this->form_add_submit_alias = 'news/add';
		$this->form_edit_submit_alias = 'news/edit';
		
		$this->output->enable_profiler(TRUE);
	}
	
	public function add_article()
	{
		$validation_result = NULL;
		$db_result = NULL;
		
		$this->load->helper('form');
		$this->load->library('form_validation');

		$data = array();
		$data['title'] = 'News Article Submittion';
		$data['form_submit_url'] = $this->form_add_submit_alias;

		$this->form_validation->set_rules('title', 'Title', 'required');
		$this->form_validation->set_rules('content', 'Content', 'required');
		
		$validation_result = $this->form_validation->run();

		if ($validation_result === FALSE)
		{
			$this->load->view('news/add', $data);
		}
		else
		{
			$db_result = $this->news_model->insert_news();
			
			if ($db_result === FALSE)
			{
				$this->load->view('news/add_error');
			}
			else
			{
				$this->load->view('news/add_ok');
			}
		}
	}
	
	public function edit_article_by_alias($alias)
	{
		$validation_result = NULL;
		$db_result = NULL;
		
		$this->load->helper('form');
		$this->load->library('form_validation');
		
		$data = array();
		$data['title'] = 'News Article Edit';
		$data['form_submit_url'] = $this->form_edit_submit_alias . '/' . $alias;
		$data['news_item'] = $this->news_model->get_news_by_alias($alias);

		$this->form_validation->set_rules('alias', 'Alias', 'required');
		$this->form_validation->set_rules('content', 'Content', 'required');
		$this->form_validation->set_rules('id', 'ID', 'required');
		$this->form_validation->set_rules('title', 'Title', 'required');
		
		$validation_result = $this->form_validation->run();

		if ($validation_result === FALSE)
		{
			$this->load->view('news/edit', $data);
		}
		else
		{
			$db_result = $this->news_model->edit_news();
			
			if ($db_result === FALSE)
			{
				$this->load->view('news/edit_error');
			}
			else
			{
				$this->load->view('news/edit_ok');
			}
		}
	}
	
	public function edit_article_by_id($id)
	{
		$validation_result = NULL;
		$db_result = NULL;
		
		$this->load->helper('form');
		$this->load->library('form_validation');
		
		$data = array();
		$data['title'] = 'News Article Edit';
		$data['form_submit_url'] = $this->form_edit_submit_alias . '/' . $id;
		$data['news_item'] = $this->news_model->get_news_by_id($id);
		
		$this->form_validation->set_rules('title', 'Title', 'required');
		$this->form_validation->set_rules('alias', 'Alias', 'required');
		$this->form_validation->set_rules('content', 'Content', 'required');
		
		$validation_result = $this->form_validation->run();

		if ($validation_result === FALSE)
		{
			$this->load->view('news/edit', $data);
		}
		else
		{
			$db_result = $this->news_model->edit_news();
			
			if ($db_result === FALSE)
			{
				$this->load->view('news/edit_error');
			}
			else
			{
				$this->load->view('news/edit_ok');
			}
		}
	}
	
	public function index()
	{
		$this->show_articles_titles();
	}
	
	public function show_articles_titles()
	{
		$data = array();
		$data['news_items'] = $this->news_model->get_news_titles_all();
        $data['title'] = 'News Archive';
		
		$this->load->view('news/articles', $data);
	}
	
	public function show_articles_editable()
	{
		$data = array();
		$data['news_items'] = $this->news_model->get_news_titles_all();
        $data['title'] = 'Editable News';
		
		$this->load->view('news/editables', $data);
	}
	
	public function show_article_by_alias($alias)
	{
		$data = array();
		$data['news_item'] = $this->news_model->get_news_by_alias($alias);

		if (empty($data['news_item']))
		{
			show_404();
		}

		$data['title'] = 'An Article';
		
		$this->load->view('news/article', $data);
	}
	
	public function show_article_by_id($id)
	{
		$data = array();
		$data['news_item'] = $this->news_model->get_news_by_id($id);

		if (empty($data['news_item']))
		{
			show_404();
		}

		$data['title'] = 'An Article';
		
		$this->load->view('news/article', $data);
	}
}
